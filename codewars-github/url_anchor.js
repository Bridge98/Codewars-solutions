function removeUrlAnchor(url){
  return url.split('#')[0];
}

console.log(removeUrlAnchor('www.codewars.com#about')); // returns 'www.codewars.com'
console.log(removeUrlAnchor('www.codewars.com?page=1')); // returns 'www.codewars.com?page=1'
