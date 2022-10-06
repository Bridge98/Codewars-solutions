function wave(str){
  
  // Code here
  const arr = [];
  const string = str.toLowerCase();
  const pattern = /\s/;
  
  for (let i = 0, j = 1; i < string.length; i += 1) {
    if (string.slice(i, j).match(pattern)) {
      j += 1
    } else {
      arr.push(`${string.slice(0, i)}${string.slice(i, j).toUpperCase()}${string.slice(j,)}`)
      j += 1;
    }
  }
  
  return arr;
}