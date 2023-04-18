function isPositive(n){
let b = parseInt(n,10)
if (b > 0) {
    return true
} else {
    return false
}
}

function abs(n){
let b = parseFloat(n,10)
if (b < 0) {
let m = b.toString(10)
m = m.slice(1);
b = parseFloat(m,10);
} 
return b
}
console.log(abs(-1))