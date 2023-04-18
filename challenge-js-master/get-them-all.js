export function getArchitects() {
    const e = document.getElementsByTagName("a")
    const f = document.getElementsByTagName("span")
    return [[...e],[...f]]
}

export function getClassical() {
    const e = document.getElementsByClassName('classical')
    const f = document.getElementsByTagName('a')
    let n = 0
    const i = []
    while (n < f.length) {
        if (f[n] === e[n]) {
            n++
        } else {
            i.push(f[n])
            n++
        }
    }
    console.log(i)
    return [[...e],[...i]]
}

export function getActive() {
    const e = document.getElementsByClassName('active', true)
    const f = document.getElementsByClassName('classical', false)
    return [[...e],[...f]]
}

export function getBonannoPisano() {
    const e = document.getElementById('BonannoPisano')
    const f = document.getElementsByClassName('active')
    let n = 0
    const i = []
    while (n < f.length) {
        if (f[n] === e) {
            n++
        } else {
            i.push(f[n])
            n++
        }
    }
    console.log(e)
    return [e,[...i]]
}