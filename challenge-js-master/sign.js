function sign(a){
let c = parseFloat(a,10)
if (c > 0) {
return 1
} else if (c === 0) {        
return 0
} else {
    return -1
}
}

function sameSign(a,b){
    let c = parseFloat(a,10)
    let d = parseFloat(b,10)
    if (sign(c) === sign(d)) {
        return true
    } else {
        return false
    }
}
console.log(sign(0))