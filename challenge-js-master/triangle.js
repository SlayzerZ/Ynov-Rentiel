function triangle(s,h) {
    let lineToPrint = "";
    for (let i = 1; i <= h; i++) {
        // for(let j = h - i; j > 0; j--) { // prepare spaces
        //   lineToPrint += " ";
        // }  
        for(let j = i * 1; j > 0; j--) { // prepare triangle symbols
          lineToPrint += s;
        }
        if (i < h) {
        lineToPrint += "\n";
      }}
      return lineToPrint;
}
console.log(triangle("b",5))