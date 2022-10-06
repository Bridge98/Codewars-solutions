#include <stdlib.h>

long long *tribonacci(const long long signature[3], size_t n) {

   long long *result = malloc(n*sizeof(long long));
  
    if(n==0) return NULL;
    else {
      for(int i=0; i<3; ++i) result[i] = signature[i];
      for(int i=3; i<n; ++i) result[i] = result[i-3]+result[i-2]+result[i-1];
    }
    
  return result;
}