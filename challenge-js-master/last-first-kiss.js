function first(n){
const b = n[0];
return b
}
    
function last(n){
const b = n[n.length - 1]
return b
}
    
function kiss(n){
const b = Array(n[n.length -1], n[0])
return b
}
console.log(kiss("mario"))