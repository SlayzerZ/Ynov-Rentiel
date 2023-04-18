function fToC(fahrenheit) 
{
  let fTemp = fahrenheit;
  let fToCel = (fTemp - 32) * 5 / 9;
  return fToCel
} 
function citiesOnly(arrObj){
const otherArray = arrObj.map((x) => {return x.city})
return otherArray
}
function upperCasingStates(arrObj){
    const spaces = ' '
    let e = arrObj.map((x) => {
        let i =/\s/
        let res = x.match(i)
        if (res != null){
            let z = x.split(' ')
            for (var q = 0; q < z.length; q++){
                z[q] = z[q].charAt(0).toUpperCase()+z[q].slice(1)
            }
            for (var q = 0; q < z.length; q +=2){
                let n = z[q] + spaces + z[q+1]
                return n
            }
        } else {
            return x.charAt(0).toUpperCase()+x.slice(1)
        }
    })
    return e
}
function fahrenheitToCelsius(arrObj){
    let i = arrObj.map(x => {return (Math.floor(fToC(parseFloat(x)))+'°C')})
    return i
}

function trimTemp(states){
    let array = states.map(temp=>{
        const regex = /\s/g
        const resRegex = new RegExp(regex)
        let strin = temp.temperature
        let str = strin.match(resRegex)
        let res = strin.replace(resRegex,'')
        temp.temperature = res
        return temp
    })
    return array
}

function tempForecasts(arrObj){
    let array = arrObj.map(all =>{
        const regex = /\s/g
        const resRegex = new RegExp(regex)
        const regexSpace = /\s\w/g
        const regexSp = new RegExp(regexSpace)
        let strin = all.temperature
        let res2 = all.state.match(regexSp)
        let str = strin.match(resRegex)
        // let res3 = all.state.match(regexSp)
        let res = strin.replace(resRegex,'')
        if(res2!== null){
            all.state = all.state.replace(regexSp,res2[0].toUpperCase())
        }
        res = parseInt(res)
        let cel = (res -32)*5/9
        cel = Math.floor(cel)
        let strRes = cel+'°Celsius in '+all.city+', ' + all.state[0].toUpperCase()+all.state.slice(1)
        return strRes
    })  
    return array
}
console.log(tempForecasts([
    {
      city: 'Pasadena',
      temperature: ' 101 °F',
      state: 'california',
      region: 'West',
    },
  ]))