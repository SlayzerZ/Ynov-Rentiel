let x = 400;
let y = 250;
let Top = false;
let Left = false;
let dvd = document.querySelectorAll('#dvd');
function generateColor(){
    let color = '#'
    return color += Math.random().toString(16).slice(2, 8).toUpperCase();
}
function moov(){
    let height = window.innerHeight -200;
    let width = window.innerWidth -300;
    if (y == height){
        Top = true;
        dvd[1].style.fill = generateColor();
    } else if (y == -80){
        Top = false;
        dvd[1].style.fill = generateColor();
    }
    if (Top){
        y -= 1
    } else {
        y += 1
    }
    if (x == width){
        Left = true;
        dvd[1].style.fill = generateColor();
    } else if (x == 0){
        Left = false;
        dvd[1].style.fill = generateColor();
    }
    if (Left){
        x -= 1
    } else {
        x += 1
    }
    dvd[1].style.top = `${y}px`;
    dvd[1].style.left = `${x}px`;
    window.requestAnimationFrame(moov)
}
window.requestAnimationFrame(moov)