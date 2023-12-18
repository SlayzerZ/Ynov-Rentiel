function read_prop(obj, prop) {
    return obj[prop];
}
let entryData = ""
let outputData = ""
let buttons = document.querySelectorAll('button')
let devise_entrée = 0;
let quantité = document.getElementById('entry');
let devise_sortie = 0;
let final = document.getElementById('result');
const money = {
    dollar: 0.84,
    euro: 1,
    livre: 1.14,
    bitcoin: 40000
}
function Entry(){
    let entry = [buttons[0],buttons[1],buttons[2],buttons[3]]
    for (let i = 0; i < entry.length; i++){
    entry[i].addEventListener('click', function(e) {
        let entry2 = entry.slice()
        entry2.splice(i,1)
        entry[i].style.backgroundColor = "rgb(0,90,0)"
        for (let j = 0; j < entry2.length; j++){
            entry2[j].style.backgroundColor = "gray"
        }
        e = e || window.event;
        var target = e.target || e.srcElement,
            entryData = target.dataset.devise;
            // console.log(entryData)
            if(entryData == "dollar"){
                devise_entrée = read_prop(money,"dollar")
            }
            if(entryData == "euro"){
                devise_entrée = read_prop(money,"euro")
            }
            if(entryData == "livre"){
                devise_entrée = read_prop(money,"livre")
            }
            if(entryData == "bitcoin"){
                devise_entrée = read_prop(money,"bitcoin")
            }
        }, false);
    }
    buttons[0].addEventListener('click',function(){
        final.value = (devise_entrée * quantité.value) / devise_sortie;
        if (isNaN(final.value) || final.value == Infinity){
            final.value = 0
        }
    })
    buttons[1].addEventListener('click',function(){
        final.value = (devise_entrée * quantité.value) / devise_sortie;
        if (isNaN(final.value) || final.value == Infinity){
            final.value = 0
        }
    })
    buttons[2].addEventListener('click',function(){
        final.value = (devise_entrée * quantité.value) / devise_sortie;
        if (isNaN(final.value) || final.value == Infinity){
            final.value = 0
        }
    })
    buttons[3].addEventListener('click',function(){
        final.value = (devise_entrée * quantité.value) / devise_sortie;
        if (isNaN(final.value) || final.value == Infinity){
            final.value = 0
        }
    })
}
function Output(){
    let output = [buttons[4],buttons[5],buttons[6],buttons[7]]
    for (let i = 0; i <output.length; i++){
    output[i].addEventListener('click', function(e) {
        let output2 = output.slice()
        output2.splice(i,1)
        output[i].style.backgroundColor = "rgb(0,90,0)"
        for (let j = 0; j < output2.length; j++){
            output2[j].style.backgroundColor = "gray"
        }
        e = e || window.event;
        var target = e.target || e.srcElement,
            outputData = target.dataset.devise;
            if(outputData == "dollar"){
                devise_sortie = read_prop(money,"dollar")
            }
            if(outputData == "euro"){
                devise_sortie = read_prop(money,"euro")
            }
            if(outputData == "livre"){
                devise_sortie = read_prop(money,"livre")
            }
            if(outputData == "bitcoin"){
                devise_sortie = read_prop(money,"bitcoin")
            }
        }, false);
    }
    buttons[4].addEventListener('click',function(){
        final.value = (devise_entrée * quantité.value) / devise_sortie;
        if (isNaN(final.value) || final.value == Infinity){
            final.value = 0
        }
    })
    buttons[5].addEventListener('click',function(){
        final.value = (devise_entrée * quantité.value) / devise_sortie;
        if (isNaN(final.value) || final.value == Infinity){
            final.value = 0
        }
    })
    buttons[6].addEventListener('click',function(){
        final.value = (devise_entrée * quantité.value) / devise_sortie;
        if (isNaN(final.value) || final.value == Infinity){
            final.value = 0
        }
    })
    buttons[7].addEventListener('click',function(){
        final.value = (devise_entrée * quantité.value) / devise_sortie;
        if (isNaN(final.value) || final.value == Infinity){
            final.value = 0
        }
    })
}
quantité.addEventListener('change',function(){
    quantité.value = this.value;
    final.value = (devise_entrée * quantité.value) / devise_sortie;
    if (isNaN(final.value) || final.value == Infinity){
        final.value = 0
    }
})
