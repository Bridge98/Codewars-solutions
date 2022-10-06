function isIsogram(str) {
  
  let i, j;
  let res; // resultant string from var;
  // lowercase every char of str:
  res = str.toLowerCase();
  
  for (i = 0; i < str.length; i++) {
    for (j = i+1; j < str.length; j++) {
      
      // check that we have no repeated chars in str:
      if (res[i] === res[j]) {
        return false; 
      }
    }
  }
  
  return true;
}