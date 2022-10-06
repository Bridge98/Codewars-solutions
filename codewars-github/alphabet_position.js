// replace all the char with their english alphabet position

function alphabetPosition(text) {
	
	let res = "";

	for (var i = 0; i < text.length; i++) {
		let char = text.toUpperCase().charCodeAt(i);
		if (char > 64 && char < 91) res += (char-64)+" "; // = char - first.char.alphabet + 1
	  else continue;
  }

	return res.slice(0, res.length-1);
}