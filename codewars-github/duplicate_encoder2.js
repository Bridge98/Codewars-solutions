function duplicateEncode(word) {
	
	let res = ""; // reseultant string
	word = word.toLowerCase(); // lowercase all the char of word

	for (var i = 0; i < word.length; i++) {
		if (word.indexOf(word[i]) === word.lastIndexOf(word[i])) {
			res += "(";
		} else {
			res += ")";
		}
	}

	return res;
}
