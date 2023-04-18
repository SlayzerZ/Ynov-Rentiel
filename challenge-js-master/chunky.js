function chunk(array,size){
    const perChunk = size // items per chunk    

const inputArray = array

const result = inputArray.reduce((resultArray, item, index) => { 
  const chunkIndex = Math.floor(index/perChunk)

  if(!resultArray[chunkIndex]) {
    resultArray[chunkIndex] = [] // start a new chunk
  }

  resultArray[chunkIndex].push(item)

  return resultArray
}, [])
return result
}
console.log(chunk(["a","b","c","d"],2))