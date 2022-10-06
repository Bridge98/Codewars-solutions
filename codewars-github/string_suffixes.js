function compareString(s1,s2) {
  let count = 0, min_length = Math.min(s1.length,s2.length);
  
  for (var i = 0; i < min_length; i++) {
    if (s1[i] == s2[i]) count++;
    else return count;
  }
  
  return count;
}

function stringSuffix(s) {  
  let sum = 0;
  
  for (var i = 0; i < s.length; i++) {
    sum += compareString(s,s.slice(i));
  } return sum;
}