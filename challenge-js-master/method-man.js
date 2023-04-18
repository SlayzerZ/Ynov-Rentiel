function words(n) {
let b = n.toString(10);
const m = b.split(' ');
return m
}
function sentence(n) {
const m = n.join(' ');
const b = m.toString(10);
return b
}
function yell(n) {
let b = n.toString(10);
const m = b.toUpperCase();
return m
}
function whisper(n) {
let b = n.toString(10);
const m = "*" + b.toLowerCase() + "*";
return m
}
function capitalize(n) {
let b = n.toString(10);
b = b.toLowerCase();
let m = b.charAt(0);
m = m.toUpperCase();
const f = m + b.slice(1);
return f
}
console.log(capitalize("MARIO"))