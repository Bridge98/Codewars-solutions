#include <stdlib.h>
#include <stdio.h>

int solution(int number) {

	int i, sum = 0;

	// find all the multiples of a number;
	for(int i = 1; i < number; i++) {
	   int value1 = i%5 == 0?5:0;
	   int value2 = i%3 == 0?3:0;
	   sum += (value1 + value2) >0 ?i:0;
	}

	return sum;
}