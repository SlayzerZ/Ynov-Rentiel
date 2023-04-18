function nasa(n){
    let limit = 1
    let t = []
    while (limit <= n) {
        let e = limit.toString()
        if (limit % 3 === 0 && limit % 5 !== 0) {
            e = 'NA'
        } else if (limit % 5 === 0 && limit % 3 !== 0) {
            e = 'SA'
        } else if (limit % 3 === 0 && limit % 5 === 0) {
            e = 'NASA'
        }
        t.push(e)
        limit += 1
    }
    return t.join(' ')
}
console.log(nasa(15))