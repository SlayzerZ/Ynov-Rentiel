function repeat(a,b) {
    let c = a.toString(10);
    let d = parseInt(b,10)
    let e = 1
    let c2 = c
    if (d === 0) {
        return ''
    }
    while (e < d){
        e = e + 1
        c2 = c2 + c
    }
    return c2
}
console.log(repeat("Salut", 0))