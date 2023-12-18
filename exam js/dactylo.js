let Phrase1 = "mère père voleur assassin barjo criminel mentor ordinateur calculette programme science vie mort dèces courage antipathie"
let Array1 = Phrase1.split(" ")
let Phrase2 = "adulte lâche primaire hanté ambigue lunettes peureux timide cyber aimer jouer louer inventer supprimer concéder tacler"
let Array2 = Phrase2.split(" ")
let Phrase3 = "manger-bouger.fr naruto sonic tg liberté prison passion mer pute enfoiré nique salaud part-dieu écologie précieux étrangler"
let Array3 = Phrase3.split(" ")
let Phrase4 = "mario luigi peach bowser yoshi donkey daisy rosalina luma birdo boo bob-omb koopa maskass cheep-cheep nintendo"
let Array4 = Phrase4.split(" ")
let arr;
let W = document.getElementsByClassName('W')
let Word = document.getElementById('word')
let timer = document.getElementById('timer')
let text = document.getElementById('text')
let start = document.getElementById('start')
let minutes = 1
let seconds = 0
let next = 0
let next2 = 1
let point = 0
let total = 0
let relance = true
let inter;
// Word.innerHTML = `<p class="W">${Array1}</p>`
function WordStart(array){
    Word.innerHTML = ``
    for(let i = 0; i < array.length; i++){
        let p = document.createElement('p')
        p.className = "W"
        p.innerHTML = `${array[i]}`
        Word.appendChild(p)
    }
}
function change(i){
    if (i == 1){
        arr = Array1
    }
    if (i == 2){
        arr = Array2
    }
    if (i == 3){
        arr = Array3
    }
    if (i == 4){
        arr = Array4
    }
}
function Timer(){
        if (relance){
            relance = false
            minutes = 1
            seconds = 0
        }
    if (seconds == 0){
        seconds = 59
        minutes -=1
    } else {
        seconds -= 1
    }
    timer.innerHTML = `${minutes}:${seconds}`
    if (minutes == 0 && seconds == 0){
        // clearInterval(inter)
        text.toggleAttribute('disabled')
        alert(`Your score is : ${point}/${total}`)
    }
}

start.addEventListener('click', () => {
    relance = true
    minutes = 1
    seconds = 0
    change(1)
    WordStart(arr)
    text.toggleAttribute('disabled')
    inter = setInterval(Timer,1000)
})
// document.addEventListener('change', () =>{
//     if (minutes <= 0 && seconds <= 0){
//         clearInterval(inter)
//     }
// })
document.addEventListener('keydown', (event) => {
    if (event.key == "Enter"){
        if (text.value == arr[next]){
            W[next].style.backgroundColor = 'green'
            point ++
        } else{
            W[next].style.backgroundColor = 'red'
        }
        if (next +1 == arr.length){
            next = 0
            next2++
            change(next2)
            WordStart(arr)
        } else {
            next++
        }
        total++
        text.value = ""
    }
})
let x = setInterval(function() {
    if (minutes == 0 && seconds == 0){
        clearInterval(inter)
        // alert(`Your score is : ${point}/${total}`)
    }
},1)
console.log(W)