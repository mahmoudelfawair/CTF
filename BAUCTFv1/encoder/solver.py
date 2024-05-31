
from base64 import b64decode
from Crypto.Cipher import DES

alpha='{}_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'

def pad(text):
    padding_length = 8 - (len(text) % 8)
    return text + bytes([0] * padding_length)

def prepare_key(k):
    key = k + k[::-1] 
    key = b''.join([hex(int(i))[2:].encode() for i in key])
    return key

def decrypt(key, enc):
    #padded_plaintext = pad(plaintext)
    cipher = DES.new(key, DES.MODE_ECB)
    plain_text = cipher.decrypt(enc)
    return plain_text


def solver():
   for i in range(1,9999):
       key = str(i).zfill(4)
       decrypted = decrypt(prepare_key(key), b64decode("PkVwkzpuMfsUPG8QTXbJU06JnOGWt9RjjfIcuXVx4rZT+lVt+VRHw+CwOv2fOmfH4E+xub1TdkUdaVxgJMqw+A=="))
       s2 = ''.join([chr(((i - 5) ^ 3) & 0xff) for i in decrypted])
       s1 = ''.join([alpha[(alpha.find(i) + 9) % len(alpha)] for i in s2])
       print(f"[#] key - {key} [#]\r")
       if 'BAU' in s1:
           print(s1, key)
           break;

solver()
