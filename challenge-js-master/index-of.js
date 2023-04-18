function indexOf(a,b,c) {
    let d = c;
    let e = 0;
    if (c === undefined) {
        d = 0
    }
    while (d < a.length){
        if (a[d] === b) {
            return d
        } else {
            d = d + 1
            e = e + 1
        }
    }
    return -1
}

function lastIndexOf(a,b,c){
    let d = c;
    let e = 0;
    let f = 0;
    if (c === undefined) {
        d = 0
    }
    while (e < a.length){
        if (a[d] === b) {
            f = d
            if (a[d +1] === b) {
                f = e
            }
        }
        d = d + 1
        e = e + 1
        if (d === a.length) {
            d = 0
        }
    }
    if (a[f] === b){
        return f
    } else {
        return -1
    }
}

function includes(a,b,c){
    let d = c;
    let e = 0;
    if (c === undefined) {
        d = 0
    }
    while (e < a.length){
        if (a[d] === b) {
            return true
        } else {
            d = d + 1
            e = e + 1
        }
        if (d === a.length) {
            d = 0
        }
    }
    return false
}
// console.log(lastIndexOf(["kazu", "meili", "zineb", "tommy","kazu"],"kazu",1))
console.log(lastIndexOf([0, 0, "t","t"],"t",3))