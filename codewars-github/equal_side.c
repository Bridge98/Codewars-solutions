#include <stdio.h>
#include <stdlib.h>

int find_even_index(const int *values, int length) {
	
	int i, SumL = 0, SumR = 0;

	for (i = 0; i < length; i++) {
		SumR += values[i]; // sum all the values of the array;
	}

	for (i = 0; i < length; i++) {
		
		SumR -= values[i];

		if (SumL == SumR) {
			return i;
		}

		SumL += values[i];
	}

	return -1;
}