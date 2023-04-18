function isValid(date) {
    let z = new Date(date)
    if (typeof date === 'string') {
        return false
    }
    if (z.toString() !== 'Invalid Date') {
        return true
    } else {
        return false
    }
}

function isAfter(date1,date2) {
    if (isValid(date1) && isValid(date2)) {
        let d1 = new Date(date1).getTime()
        let d2 = new Date(date2).getTime()
        if (d1 < d2) {
            return false;
          } else if (d1 > d2) {
            return true;
          } else {
            console.log(`Both dates are equal`);
          }
    }
}

function isBefore(date1,date2) {
    if (isValid(date1) && isValid(date2)) {
        let d1 = new Date(date1).getTime()
        let d2 = new Date(date2).getTime()
        if (d1 > d2) {
            return false;
          } else if (d1 < d2) {
            return true;
          } else {
            console.log(`Both dates are equal`);
          }
    }
}

function isFuture(date) {
    if (isValid(date)) {
        let d1 = new Date(date).getTime()
        let d2 = Date.now()
        if (d1 < d2) {
            return false;
          } else if (d1 > d2) {
            return true;
          } else {
            console.log(`Both dates are equal`);
          }
    }
}

function isPast(date) {
    if (isValid(date)) {
        let d1 = new Date(date).getTime()
        let d2 = Date.now()
        if (d1 > d2) {
            return false;
          } else if (d1 < d2) {
            return true;
          } else {
            console.log(`Both dates are equal`);
          }
    }
}
console.log(isValid('2013'))