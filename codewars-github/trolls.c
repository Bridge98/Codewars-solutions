//solution must allocate all required memory
//and return a free-able buffer to the caller.

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *disemvowel(const char *str) {

	int j = 0;
	char *res = malloc(sizeof(char)*strlen(str)+1);

	for (int i = 0; i < strlen(str); i++) {
		switch (str[i]) { // this remove all the vowels from a string line;
			case 'a':
			case 'e':
			case 'i':
			case 'o':
			case 'u':
			case 'A':
			case 'E':
			case 'I':
			case 'O':
			case 'U':
				break;
			default:
				res[j++] = str[i];
		}
	}
  
  res[j] = '\0';
	return res;
}