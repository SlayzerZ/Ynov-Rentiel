import { styles } from './pimp-my-style.data.js'
export function pimp(){
    const button = document.getElementsByTagName('button')
    let m = button[0].className.split(' ')
    let ze = m.length -1
    if (m.length === 0) {
        button[0].className = 'button'
    }
    let pimp = true
    if (m.length === 3 && pimp === false){
        pimp = true
        m.splice(2,1)
        button[0].className = m.join(' ')
    }
    if (ze === styles.length) {
        if (m[ze] === styles[styles.length -1]) {
        button[0].classList.add("unpimp")
        m.push('unpimp')
        }
        pimp = false
}
    if (m[ze] === 'unpimp' && pimp === true) {
        pimp = false
    }
    if (pimp === true) {
        button[0].classList.add(styles[ze]) 
        ze += 1
        if (ze === styles.length) {
            button[0].classList.add("unpimp")
        }
    } else if (pimp === false) {
        if (m[ze] === 'unpimp') {
            m.splice(ze-1,1)
            if (m.length === 2) {
                m.splice(1,1)
            }
        }
        button[0].className = m.join(' ')
        ze = ze - 1
    }
    // console.log(m)
}