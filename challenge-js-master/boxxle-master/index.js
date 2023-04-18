import { Levels } from './level.js'

function nestedCopy(array) {
    return JSON.parse(JSON.stringify(array));
}
let audio1 = document.getElementById('audio1');
let collision = document.getElementById('collision');
let soundbox = document.getElementById('soundbox');
let item = document.getElementById('item');
let MarioYell = document.getElementById('Mario');
let Course = document.getElementById('Course');
let Fgame = document.getElementById('Fgame');
let Game = document.getElementById('Game');
let star = false
let firstGame = false;
let playerPosition = false;
function toggleMusic(x){
    if(x.className === "play"){
        x.pause();
        x.className = "mute"
    }
    if(x.className === "mute"){
        x.play();
        x.className = "play"
    }
}

let startBox = 0
let startBox2 = 0
let sBox = 0 
let fBox = 0
let stepCounter = 0
let a = 0
let b = 0
let c = 0
let x = 0
let y = 0

function nextLevel(){
    if (a+1 <= Levels.length){
        a += 1
        if (Levels[a] === undefined){
            Fgame.play()
            // Options.appendChild(Bowser)
            // Game.play()
            // alert("T'as gagné, t'es trop fort!!!")
        } else {
            reload()
        }
    }
    
}
function previousLevel(){
    if (a-1 >= 0){
        a -= 1
        reload()
    }
}

let actual = nestedCopy(Levels[a])

let player1 = new Image (50,50); player1.src = './assets/pics/mario1.gif'
let player2 = new Image (50,50); player2.src = './assets/pics/mario2.gif'
let player3 = new Image (50,50); player3.src = './assets/pics/mario3.gif'
let player4 = new Image (50,50); player4.src = './assets/pics/mario4.gif'
let player5 = new Image (50,50); player5.src = './assets/pics/mario5.gif'
// let Bowser = new Image (200,150); Bowser.src = './assets/pics/fuck-you-bowser.gif'
// let audio1i = new Image (50,50); audio1i.src = './assets/pics/play.jpg'

// if(audio1.className === "play"){
//     audio1i.src = './assets/pics/play.jpg'
// }
// if(audio1.className === "mute"){
//     audio1i.src = './assets/pics/mute.jpg'
// }

 let para = document.createElement('div')
 para.id = "Step-counter"
 
 let name = ""
 let Options = document.getElementById('Options')
 let canvas;
 let context;
 
 let player = player1
//  Options.appendChild(audio1i)
 
//  audio1i.addEventListener('click', toggleMusic(audio1))
 
 window.onload = init;

 function reload(){
    star = false
    context.remove()
    actual = nestedCopy(Levels[a])
    playerPosition = false
    firstGame = false
    finish(actual);
    if(playerPosition){
        drawGame();
}
    stepCounter = 0
 }

 function init(){
    
     canvas = document.getElementById('Box');
     // Start the first frame request
     finish(actual);
     if(playerPosition){
     drawGame();
     MarioYell.play();
    }
    let R = document.getElementById('reload');
    R.addEventListener('click', reload)
    
    audio1.className = "mute"
    
    window.requestAnimationFrame(gameLoop);
 }
 
 function gameLoop(){
    if(!playerPosition){
        alert("The game can't start")
    } else {
    if (name === "ArrowDown"){
        stepCounter += 1
        para.innerHTML = "Nombre de pas :" + stepCounter
        Options.appendChild(para)
        player = player1
        name = ""
        context.remove()
        if (actual[y+1][x] === 0) {
            if (star){
                actual[y+1][x] = 3
                actual[y][x] = 4
                star = false
            } else {
                actual[y+1][x] = 3
                actual[y][x] = 0
                
            }
        }
        if (actual[y+1][x] === 4) {
            item.play()
            if (star){
                actual[y+1][x] = 3
                actual[y][x] = 4
            } else {
                actual[y+1][x] = 3
                actual[y][x] = 0
                star = true
            } 
        }
        if (actual[y+1][x] === 2) {
            soundbox.play()
           if (actual[y+2][x] === 4) {
            fBox += 1
            sBox -= 1
            actual[y+2][x] = 5
            actual[y+1][x] = 3
            actual[y][x] = 0
            }else if (actual[y+2][x] === 0) {
                actual[y+2][x] = 2
                actual[y+1][x] = 3
                actual[y][x] = 0
            }
        }
        if (actual[y+1][x] === 5) {
            soundbox.play()
            if (actual[y+2][x] === 4) {
                if (star){
                    actual[y+2][x] = 5
                    actual[y+1][x] = 3
                    actual[y][x] = 4
                } else {
                    actual[y+2][x] = 5
                    actual[y+1][x] = 3
                    actual[y][x] = 0
                    star = true
                }
             
             }else if (actual[y+2][x] === 0) {
                fBox -= 1
                sBox += 1
                 actual[y+2][x] = 2
                 actual[y+1][x] = 3
                 actual[y][x] = 4
                 star = true
             }
         }
         if (actual[y+1][x] === 1){
            collision.play()
         }
        drawGame()
    }
    if (name === "ArrowLeft"){
        stepCounter += 1
        para.innerHTML = "Nombre de pas :" + stepCounter
        Options.appendChild(para)
        player = player2
        name = ""
        context.remove()
        if (actual[y][x-1] === 0) {
            if (star){
                actual[y][x-1] = 3
                actual[y][x] = 4
                star = false
            } else {
                actual[y][x-1] = 3
                actual[y][x] = 0
                
            }
        }
        if (actual[y][x-1] === 4) {
            item.play()
            if (star){
                actual[y][x-1] = 3
                actual[y][x] = 4
            } else {
                actual[y][x-1] = 3
                actual[y][x] = 0
                star = true
            } 
        }
        if (actual[y][x-1] === 2) {
            soundbox.play()
           if (actual[y][x-2] === 4) {
            fBox += 1
            sBox -= 1
            actual[y][x-2] = 5
            actual[y][x-1] = 3
            actual[y][x] = 0
            }else if (actual[y][x-2] === 0) {
                if (star){
                    actual[y][x-2] = 2
                    actual[y][x-1] = 3
                    actual[y][x] = 4
                    star = false
                } else {
                    actual[y][x-2] = 2
                    actual[y][x-1] = 3
                    actual[y][x] = 0
                }
            }
        }
        if (actual[y][x-1] === 5) {
            soundbox.play()
            if (actual[y][x-2] === 4) {
                if (star){
                    actual[y][x-2] = 5
                    actual[y][x-1] = 3
                    actual[y][x] = 4
                } else {
                    actual[y][x-2] = 5
                    actual[y][x-1] = 3
                    actual[y][x] = 0
                    star = true
                }
             
             }else if (actual[y][x-2] === 0) {
                fBox -= 1
                sBox += 1
                actual[y][x-2] = 2
                 actual[y][x-1] = 3
                 actual[y][x] = 0
                 star = true
             }
         }
         if (actual[y][x-1] === 1){
            collision.play()
         }
        drawGame()
    }
    if (name === "ArrowRight"){
        stepCounter += 1
        para.innerHTML = "Nombre de pas :" + stepCounter
        Options.appendChild(para)
        player = player3
        name = ""
        context.remove()
        if (actual[y][x+1] === 0) {
            if (star){
                actual[y][x+1] = 3
                actual[y][x] = 4
                star = false
            } else {
                actual[y][x+1] = 3
                actual[y][x] = 0
                
            }
        }
        if (actual[y][x+1] === 4) {
            item.play()
            if (star){
                actual[y][x+1] = 3
                actual[y][x] = 4
            } else {
                actual[y][x+1] = 3
                actual[y][x] = 0
                star = true
            } 
        }
        if (actual[y][x+1] === 2) {
            soundbox.play()
           if (actual[y][x+2] === 4) {
            fBox += 1
            sBox -= 1
            actual[y][x+2] = 5
            actual[y][x+1] = 3
            actual[y][x] = 0
            }else if (actual[y][x+2] === 0) {
                if (star){
                    actual[y][x+2] = 2
                    actual[y][x+1] = 3
                    actual[y][x] = 4
                    star = false
                } else {
                    actual[y][x+2] = 2
                    actual[y][x+1] = 3
                    actual[y][x] = 0
                }
            }
        }
        if (actual[y][x+1] === 5) {
            soundbox.play()
            if (actual[y][x+2] === 4) {
                if (star){
                    actual[y][x+2] = 5
                    actual[y][x+1] = 3
                    actual[y][x] = 4
                } else {
                    actual[y][x+2] = 5
                    actual[y][x+1] = 3
                    actual[y][x] = 0
                    star = true
                }
             
             }else if (actual[y][x+2] === 0) {
                fBox -= 1
                sBox += 1
                 actual[y][x+2] = 2
                 actual[y][x+1] = 3
                 actual[y][x] = 0
                 star = true
             }
         }
         if (actual[y][x+1] === 1){
            collision.play()
         }
        drawGame()
    }
    if (name === "ArrowUp"){
        stepCounter += 1
        para.innerHTML = "Nombre de pas :" + stepCounter
        Options.appendChild(para)
        player = player4
        name = ""
        context.remove()
        if (actual[y-1][x] === 0) {
            if (star){
                actual[y-1][x] = 3
                actual[y][x] = 4
                star = false
            } else {
                actual[y-1][x] = 3
                actual[y][x] = 0
                
            }
        }
        if (actual[y-1][x] === 4) {
            item.play()
            if (star){
                actual[y-1][x] = 3
                actual[y][x] = 4
            } else {
                actual[y-1][x] = 3
                actual[y][x] = 0
                star = true
            }
        }
        if (actual[y-1][x] === 2) {
            soundbox.play()
           if (actual[y-2][x] === 4) {
            fBox += 1
            sBox -= 1
            actual[y-2][x] = 5
            actual[y-1][x] = 3
            actual[y][x] = 0
            }else if (actual[y-2][x] === 0) {
                actual[y-2][x] = 2
                actual[y-1][x] = 3
                actual[y][x] = 0
            }
        }
        if (actual[y-1][x] === 5) {
            soundbox.play()
            if (actual[y-2][x] === 4) {
                if (star){
                    actual[y-2][x] = 5
                    actual[y-1][x] = 3
                    actual[y][x] = 4
                } else {
                    actual[y-2][x] = 5
                    actual[y-1][x] = 3
                    actual[y][x] = 0
                    star = true
                }
             
             }else if (actual[y-2][x] === 0) {
                fBox -= 1
                sBox += 1
                 actual[y-2][x] = 2
                 actual[y-1][x] = 3
                 actual[y][x] =4
                 star = true
             }
         }
         if (actual[y-1][x] === 1){
            collision.play()
         }
        drawGame()
    }
    if (fBox === startBox2 && sBox === 0){
        if(Levels[a+1] !== undefined){
            Course.play()
        } 
        nextLevel()
    }
    if (name === "r" || name === "R"){
        reload()
    }
    
     window.requestAnimationFrame(gameLoop);
 }
}

 const drawGame = () => {
    let table1 = document.createElement("table")
    b = 0
    while (b < actual.length){
        for (c = 0;c < actual[b].length;c++) {
            let line = document.createElement("tr")
            let box = new Image (50,50); box.src = './assets/pics/box.png'
            let box2 = new Image (50,50); box2.src = './assets/pics/box2.png'
            let wall = new Image (50,50); wall.src = './assets/pics/wall.jpg'
            let goal = new Image (50,50); goal.src = './assets/pics/star.png'
            if (actual[b][c] === 0){
                line.className = "vide"
            }
            if (actual[b][c] === 1){
                line.className = "wall"
                line.appendChild(wall)
            }
            if (actual[b][c] === 2){
                line.className = "box"
                line.appendChild(box)
                // startBox += 1
            }
            if (actual[b][c] === 3){
                line.className = "player"
                x = nestedCopy(c)
                y = nestedCopy(b)
                line.appendChild(player)
            }
            if (actual[b][c] === 4){
                line.className = "goal"
                line.appendChild(goal)
            }
            if (actual[b][c] === 5){
                line.className = "box2"
                line.appendChild(box2)
            }
            table1.appendChild(line)
        }
        b++
    }
    if (!firstGame) {
        startBox2 = nestedCopy(startBox)
        sBox = nestedCopy(startBox)
        firstGame = true
    }
    table1.id = "gameboard"
    context = table1
    canvas.appendChild(table1)
    
 }
 
 function finish(x){
    startBox = 0
    b = 0
    while (b < x.length){
        for (c = 0;c < x[b].length;c++) {
            if (x[b][c] === 2){
                startBox += 1
            }
            if (x[b][c] === 3){
                playerPosition = true
            }
        }
        b++
    }
    if (!firstGame) {
        startBox2 = nestedCopy(startBox)
        sBox = nestedCopy(startBox)
        fBox = 0
        firstGame = true
    }
    
 }
 
 document.addEventListener('keydown', (event) => {
    name = event.key;
    if (name === "n" || name === "N"){
        nextLevel()
    }
    if (name === "p" || name === "P"){
        previousLevel()
    }
})
