function multiply(a, b) {
let c = parseFloat(a,10)
let d = parseFloat(b,10)
let e = 2
let f = c
if (d < 0) {
    let d2 = d.toString();
    d2 = d2.slice(1);
    d = parseFloat(d2,10)
}
while (e <= d) {
    e = e + 1
    f = f + c
}
if (a === 1) {
    return d
} else if (a === 0) {
return 0
} else if (b === 0){
    return 0
} else {
    if (b < 0 && a > 0) {
        let d2 = f.toString();
        d2 = '-' + d2
        f = parseFloat(d2,10)
    } else if (b < 0 && a < 0) {
        let d2 = f.toString();
        d2 = d2.slice(1);
        f = parseFloat(d2,10)
    }
    return f
}
}

function divide(a, b) {
    if (b === 0) {
    //   console.log('Division by zero is undefined: ' + a + '/' + b);
      return 0
    }
    let sign = 1;
    if (a < 0) {
      a = -a;
      sign = -sign;
    }
    if (b < 0) {
      b = -b;
      sign = -sign;
    }
    let result = 0;
    while (a >= 0) {
      a -= b;
      result++;
    //   console.log(a)
    }
    return multiply((result - 1),sign);
  }

function modulo(a, b) {
    let c = a
    if (b === 0) {
        // console.log('Division by zero is undefined: ' + a + '/' + b);
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

console.log(modulo(34,78))