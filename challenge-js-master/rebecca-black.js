function isFriday(date){
let days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"]
    let z = new Date(date)
    let a = date.getDay()
    if (a === 5) {
        return true
    } else {
        return false
    }
}
function isWeekend(date){
    let a = date.getDay()
    if (a+1 > 5) {
        return true
    } else {
        return false
    }
}
function isLeapYear(date) {
    const leap = date.getYear();
    const leap2 = new Date(leap, 1, 29).getDate() === 29;
    if (leap2) {
        return true
    } else {
        return false
    }
}
function isLastDayOfMonth(date){
    let n = 0
    let e = date.toString()
    let last_date = new Date(date.getFullYear(), date.getMonth() + 1, 0);
    let m = last_date.toString();
    let a = e.split(" ");
    let z = m.split(" ")
    e = ""
    m = ""
    while (n !== 4) {
       e += a[n] + ' '
       m += z[n] + ' '
       n += 1
    }
    console.log(m)
    if (e === m) {
        return true
    } else {
        return false
    }
}
console.log(isLastDayOfMonth(new Date('2020-02-29')))