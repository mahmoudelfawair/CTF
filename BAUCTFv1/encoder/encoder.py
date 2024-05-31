from Crypto.Cipher import DES

def pad(text):
    padding_length = 8 - (len(text) % 8)
    return text + bytes([0] * padding_length)

def prepare_key(k):
    key = k + k[::-1] 
    key = b''.join([hex(int(i))[2:].encode() for i in key])
    return key


def encrypt(key, plaintext):
    padded_plaintext = pad(plaintext)
    cipher = DES.new(key, DES.MODE_ECB)
    encrypted = cipher.encrypt(padded_plaintext)
    return encrypted

alpha = '{}_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'

flag = "BAUCTF{fake_flag}"

# tarameso remember it is 4 digits 
key = '0000'

key = prepare_key(key)

stage1 = ''.join([alpha[(alpha.find(i) - 9) % len(alpha)] for i in flag])

stage2 = b''.join([chr((ord(i) ^ 3) + 5).encode() for i in stage1])

stage3 = encrypt(key, stage2)


# with open('enc.out', 'bw') as f:
#     f.write(stage3)