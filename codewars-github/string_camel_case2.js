function toCamelCase(str){
  
  //console.log(str, 'testing')
  
  if(str === ''){
    return ''
  } else {

   let containmentArea = []
   let splitString = str.replace(/[^A-Z0-9]/ig, "_").split("_")
   //[ 'the', 'stealth', 'warrior' ]
   let firstElement = containmentArea.push( splitString.splice(0,1) )

   for(let word in splitString){

     let splitWords = splitString[word].split('')
     let capitalLetter = splitWords[0].toUpperCase()

     splitWords.splice(0,1, capitalLetter)
     let joinedWord = splitWords.join('')

     containmentArea.push(joinedWord)
     let newSentence = containmentArea.join('')

   }
     return containmentArea.join('')
  }
}