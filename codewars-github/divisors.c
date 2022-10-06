// count the divisors of a number;
#include <stdio.h>

int divisors(int n) {

	int i, count=0;

	for (i = 1; i <= n; i++) {
		if (n%i == 0)
			count++; 

	}
	
	printf("divisors(%d) == %d\n",n,count);
	return (n,count); 
}