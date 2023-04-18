function modulo(a, b) {
    let c = a
    if (b === 0) {
        return 0
      }
      let sign = 1;
      if (c < 0) {
        c = -c;
        sign = -sign;
      }
      if (b < 0) {
        b = -b;
        sign = -sign;
      }
      let result = 0;
      while (c >= 0) {
        c -= b;
        result++;
      //   console.log(a)
      }
      let f = c + b
      if (a < 0 && b > 0) {
        f = -f
    }
      return f;
}

function round(a) {
    if (a > 0) {
    let c = modulo(a,1)
    let d = a - c
    if (c <= 0.5) {
        return d
    } else {
        return d + 1
    }
} else if (a < 0) {
    a = a * -1
    let c = modulo(a,1)
    let d = a - c
    if (c <= 0.5) {
        return d * -1
    } else {
        return (d + 1) * -1
    }
} else {
    return a
}

}

function ceil(a) {
    if (a > 0) {
        let c = modulo(a,1)
        let d = a - c
        if (c === 0) {
            return d
        } else {
            return d + 1
        }
    } else if (a < 0) {
        a = a * -1
        let c = modulo(a,1)
        let d = a - c
        // console.log(d)
        if (c === 0) {
            return ((d) * -1)
        } else {
            return (d) * -1
        }
    } else {
        return a
}
}

function floor(a) {
    if (a > 0) {
        let c = modulo(a,1)
        let d = a - c
        if (c === 0) {
            return d
        } else {
            return d
        }
    } else if (a < 0) {
        a = a * -1
        let c = modulo(a,1)
        let d = a - c
        if (c === 0) {
            return d * -1
        } else {
            return (d + 1) * -1
        }
    } else {
        return a
}
}

function trunc(a) {
    let overflow = 0
    let d = 0
    let r = 0
    if (a >= 0xfffffffff) {
        a -= 0xfffffffff
        r = floor(a)
        overflow = 1
        return r + 0xfffffffff
    }
    if (a > 0) {
        let c = modulo(a,1)
        d = a - c
        return d
    } else if (a < 0) {
        a = a * -1
        let c = modulo(a,1)
        d = a - c
       return d * -1
    } if (overflow === 1) {
        let c = modulo(a,1)
        d = a - c
        return d += 0xfffffffff
    } else {
        return a
    }
}
console.log(trunc(-15.25))