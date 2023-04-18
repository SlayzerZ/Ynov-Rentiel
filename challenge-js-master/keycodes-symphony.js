function notLow(message){
    let alphabet = ["a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"," "]
    let z = 0
    while (z !== alphabet.length) {
        if (message === alphabet[z]){
            return true;
        } else {
            z += 1
        }
    }
    return false
}

export function compose(){
// Add event listener on keydown
document.addEventListener('keydown', (event) => {
    var name = event.key;
    var code = event.code;
    // alert(`Key pressed ${name} \r\n Key code value: ${code}`)
    // Alert the key name and key code on keydown
    let e = notLow(name);
    // console.log(e);
    if (name !== "Backspace") {
    if (name !== "Escape") {
    if (e) {
    const keyNode = document.createElement("div");
    if (name === "a") {
        keyNode.style.backgroundColor = '#FC4441';
    }
    if (name === "b") {
        keyNode.style.backgroundColor = '#8cc068';
    }
    if (name === "c") {
        keyNode.style.backgroundColor = '#8A3A1B';
    }
    if (name === "d") {
        keyNode.style.backgroundColor = '#EF43C8';
    }
    if (name === "e") {
        keyNode.style.backgroundColor = '#CD45C6';
    }
    if (name === "f") {
        keyNode.style.backgroundColor = '#4C05DB';
    }
    if (name === "g") {
        keyNode.style.backgroundColor = '#77028D';
    }
    if (name === "h") {
        keyNode.style.backgroundColor = '#705D3C';
    }
    if (name === "i") {
        keyNode.style.backgroundColor = '#95594A';
    }
    if (name === "j") {
        keyNode.style.backgroundColor = '#A4AD79';
    }
    if (name === "k") {
        keyNode.style.backgroundColor = '#EEBDEA';
    }
    if (name === "l") {
        keyNode.style.backgroundColor = '#D29179';
    }
    if (name === "m") {
        keyNode.style.backgroundColor = '#6CC8FC';
    }
    if (name === "n") {
        keyNode.style.backgroundColor = '#6BBAE1';
    }
    if (name === "o") {
        keyNode.style.backgroundColor = '#A321FB';
    }
    if (name === "p") {
        keyNode.style.backgroundColor = '#2CFCAE';
    }
    if (name === "q") {
        keyNode.style.backgroundColor = '#2571A8';
    }
    if (name === "r") {
        keyNode.style.backgroundColor = '#28B2DE';
    }
    if (name === "s") {
        keyNode.style.backgroundColor = '#55E90F';
    }
    if (name === "t") {
        keyNode.style.backgroundColor = '#3D8792';
    }
    if (name === "u") {
        keyNode.style.backgroundColor = '#198C1E';
    }
    if (name === "v") {
        keyNode.style.backgroundColor = '#CE15D5';
    }
    if (name === "w") {
        keyNode.style.backgroundColor = '#6479F2';
    }
    if (name === "x") {
        keyNode.style.backgroundColor = '#D21303';
    }
    if (name === "y") {
        keyNode.style.backgroundColor = '#9D6586';
    }
    if (name === "z") {
        keyNode.style.backgroundColor = '#12544D';
    }
    keyNode.id = code;
    keyNode.className = 'note'
    keyNode.innerText = name
    document.body.append(keyNode);
    }
    } else {
        const bricks = document.querySelectorAll('div')
        let z = 0
        while (z !== bricks.length) {
            bricks[z].remove()
            z += 1
        }
    }
} else if (name === "Backspace") {
        const bricks = document.querySelectorAll('div')
        bricks[bricks.length-1].remove()
    }
  }, false);
  
}

function logMessage(message) {
    consoleLog.innerHTML += `${message}<br>`;
  }