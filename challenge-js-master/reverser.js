function reverse(a){
    let b = [];
    let c = 0;
    let t = ''
    while (c < a.length){
        c = c + 1
        b.push(a[a.length-c])
        t += b[c-1]
    }
    if (Array.isArray(a)) {
    return b
    } else {
        return t
    }
}
console.log(reverse([15,56,54,89,56,54,52]))
console.log(reverse("monsieur"))