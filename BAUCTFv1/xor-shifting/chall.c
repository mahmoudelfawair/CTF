#include <stdio.h>
#include <string.h>
int main(){
	

	char flag[64];
	printf("Enter the flag: ");
	scanf("%s",flag);
	unsigned char prev_t = flag[0];
	int tmp_int; 
	char res;
	printf("%02x", prev_t);
	for(int i = 1; i < strlen(flag); i++){
		tmp_int = prev_t >> 4;
		res = flag[i] ^ tmp_int;
	  	//printf("res %d - prev_t %d - tmp_int %d\n", res, prev_t, tmp_int);     	
		printf("%02x", res);	
		prev_t = flag[i];
	}
}
