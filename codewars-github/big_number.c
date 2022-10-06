#include <stdbool.h>
#include <math.h>

bool narcissistic(int num) {
	int original, rem, sum=0, digit=0;
	original = num;

	/* Counting number of digit in a given number */
	while(num!=0)
	{
	  digit++;
	  num = num/10;
	}

	/* After execution above loop number becomes 0
	So copying original number to variable number */

	num = original;
	
	/* Finding sum */
	while(num != 0)
	{
	  rem = num%10;
	  sum = sum + pow(rem, digit);
	  num = num/10;
	}
	
	/***Check if Armstrong or not***/
	if(sum == original)
	{
		return(true);
	}
	else
	{
		return(false);
	}
	
	return(true);
}