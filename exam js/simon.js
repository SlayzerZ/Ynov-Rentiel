let sequence = []
let hsequence = []
let level = 0
let minutes = 10
let seconds = 0
let index = 0
let Buttons = document.querySelectorAll('.eightButton')
let Start = document.querySelector('.start')
let game = document.querySelector('.gameContainer')
let Count = document.getElementById('CountText')
let Title = document.getElementById('Title')
let Timer = document.getElementById('Timer')
let relance = false
let playerTime = false
let x = setInterval(function () {
    if (relance){
        relance = false
        minutes = 10
        seconds = 0
    }
if (seconds == 0){
    seconds = 59
    minutes -=1
} else {
    seconds -= 1
}
if (minutes == 0 && seconds == 0){
    resetGame('Oops! Game over, you pressed the wrong Button');
    clearInterval();
    localStorage.setItem(1,level)
    return;
}
Timer.innerHTML = `${minutes}mm ${seconds}s`
},1000)
function humanTurn(){
    game.classList.toggle('unclick')
    playerTime = true
}
function activate(button){
    button.classList.add('activated')
    button.style.boxShadow = "0 5px #666"
    button.style.transform = "translateY(4px)"
    setTimeout(() => {
        button.classList.remove('activated');
        button.style.boxShadow = "";
        button.style.transform = "";
      }, 300);
}
function playRound(nextSequence) {
    nextSequence.forEach((color, index) => {
      setTimeout(() => {
        activate(color);
      }, (index + 1) * 600);
    });
  }
function step(){
    let Buttons = document.querySelectorAll('.eightButton')
    let random = Buttons[Math.floor(Math.random() * Buttons.length)];
    return random
}
function next(){
    game.classList.toggle('unclick')
    playerTime = false
    level += 1
    let nextsec = [...sequence]
    nextsec.push(step())
    playRound(nextsec)
    sequence = [...nextsec];
    setTimeout(() => {
    humanTurn(level);
  }, level * 600 + 1000);
  console.log(sequence)
}
function resetGame(text){
    alert(text);
    sequence = [];
    hsequence = [];
    level = 0;
    Start.classList.toggle('active')
    Start.style.backgroundColor = ""
}
function resetGame2(){
    sequence = [];
    hsequence = [];
    level = 0;
}
function handle(button){
    hsequence.push(button);
    console.log(hsequence)
    // let remaintaps = sequence.length - hsequence.length;
    // Title.innerHTML = `Simon<br> You have ${remaintaps} taps left`
        if (hsequence[index] !== sequence[index]) {
            resetGame('Oops! Game over, you pressed the wrong Button');
            return;
          }
          index++;
        if (hsequence.length === sequence.length){
            next();
            hsequence = []
            index = 0
        }
}

function startGame(){
    relance = true
    sequence = []
    level = 0
    Start.classList.toggle('active')
    let con = Start.classList;
    // console.log(con)
    if (con.length > 1){
        Start.style.backgroundColor = "green"
        next();
    } else{
        Start.style.backgroundColor = ""
        resetGame2();
        // game.classList.toggle('unclick')
    }
}
Start.addEventListener('click', startGame)
Buttons[0].addEventListener('click', function() {
if (playerTime){
    handle(Buttons[0])
}
})
Buttons[1].addEventListener('click', function() {
    if (playerTime){
        handle(Buttons[1])
    }
})
Buttons[2].addEventListener('click', function() {
    if (playerTime){
        handle(Buttons[2])
    }
})
Buttons[3].addEventListener('click', function() {
    if (playerTime){
        handle(Buttons[3])
    }
})
Buttons[4].addEventListener('click', function() {
    if (playerTime){
        handle(Buttons[4])
    }
})
Buttons[5].addEventListener('click', function() {
    if (playerTime){
        handle(Buttons[5])
    }
})
Buttons[6].addEventListener('click', function() {
    if (playerTime){
        handle(Buttons[6])
    }
})
Buttons[7].addEventListener('click', function() {
    if (playerTime){
         handle(Buttons[7])
    }
})
document.addEventListener('click', function() {
Count.innerHTML = level
})