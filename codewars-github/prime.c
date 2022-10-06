#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <math.h>

bool is_prime(int num) {

	int i;

	if (num < 2) return false;
	for (i = 2; i <= sqrt(num); i++) {
		if (num%i == 0) return false;
	}

	return true;
}