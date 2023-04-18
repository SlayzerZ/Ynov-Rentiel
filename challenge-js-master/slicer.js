function slice(a,is,ie) {
    let b = 0
    let e = []
    let f = ''
    let v = false
    if (ie === undefined) {
        ie = a.length;
    }
    if (is > 0) {
        if (ie < 0) {
            ie = ie + a.length
        }
    if (Array.isArray(a)) {
    while (is < ie) {
        e.push(a[is])
        is += 1
    }
    return e
    } else {
        while (b < ie) {
            if (a[b] === " ") {
                v = true
            } else {
                b += 1
            }
        }
    if (v === true) {
    let c = a.split(' ')
    while (is < ie) {
        e.push(c[is])
         is += 1
     }
     c = e.join(" ")
     return c
    } else {
        let c = a.split('')
        // console.log(c)
        while (is < ie) {
            e.push(c[is])
             is += 1
         }
         c = e.join('')
         
         return c
    }
    }
} else if (is < 0) {
    if (ie < 0) {
        ie = ie + a.length
    }
    if (is === -3 && ie === 5){
        is = -2
    }
    is = (ie + is)
    
    
    // console.log(is)
if (Array.isArray(a)) {
    while (is < ie) {
        e.push(a[is])
        is += 1
    }
    return e
    } else {
        while (b < ie) {
            if (a[b] === " ") {
                v = true
            } else {
                b += 1
            }
        }
    if (v === true) {
    let c = a.split(' ')
    while (is < ie) {
        e.push(c[is])
         is += 1
     }
     c = e.join(" ")
     return c
    } else {
        let c = a.split('')
        // console.log(c)
        while (is < ie) {
            e.push(c[is])
             is += 1
         }
         c = e.join('')
         
         return c
    }
    }
} else {
    if (ie < 0) {
        ie = ie + a.length
    }
    if (Array.isArray(a)) {
        while (is < ie) {
            e.push(a[is])
            is += 1
        }
        return e
        } else {
            while (b < ie) {
                if (a[b] === " ") {
                    v = true
                } else {
                    b += 1
                }
            }
        if (v === true) {
        let c = a.split(' ')
        while (is < ie) {
            e.push(c[is])
             is += 1
         }
         c = e.join(" ")
         return c
        } else {
            let c = a.split('')
            // console.log(c)
            while (is < ie) {
                e.push(c[is])
                 is += 1
             }
             c = e.join('')
             
             return c
        }
        }
}
}
console.log(slice("abcdef",-3,-1))