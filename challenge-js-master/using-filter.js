function filterShortStateName(arr){
    const result = arr.filter(word => word.length < 7)
    return result
}
function filterStartVowel(arr){
    const result = arr.filter(word => word.startsWith("A")+word.startsWith("E")+word.startsWith("I")+word.startsWith("O")+word.startsWith("U"))
    return result
}
function filter5Vowels(obj){
    let array = obj.map(vowels5=>{
        const regex = /[aeiouAEIOU]/g
        const resRegex = new RegExp(regex)
        let str = vowels5.match(resRegex)
        if(str.length >= 5){
            return vowels5
        }
        return null
    })
    let arr = []
    arr = array.filter(function(s){
        return s!==null
    })
    return arr
}
function filter1DistinctVowel(obj){
    let array = obj.map(vowel=>{
        const reg = /[aeiouAEIOU]/g
        const rReg = new RegExp(reg)
        let str = vowel.match(rReg)
        // let arr = []
        let flag = false
        Loop :
        for(let i=0;i<str.length-1;i++){
            for(let j=i+1;j<str.length;j++){
                if(str[i].toUpperCase() !== str[j].toUpperCase()){
                    flag = false
                    break Loop
                }
            }
            flag = true
        }
        if(flag){
            return vowel
        }
        return null
    })
    let arr = []
    arr = array.filter(function(s){
        return s!==null
    })
    return arr
}
console.log(filter1DistinctVowel([
    'Alabama',
    'Alaska',
    'Arizona',
    'Arkansas',
    'California',
    'Colorado',
    'Connecticut',
    'Delaware',
    'Florida',
    'Georgia',
    'Hawaii',
    'Idaho',
    'Illinois',
    'Indiana',
    'Iowa',
    'Kansas',
    'Kentucky',
    'Louisiana',
    'Maine',
    'Maryland',
    'Massachusetts',
    'Michigan',
    'Minnesota',
    'Mississippi',
    'Missouri',
    'Montana',
    'Nebraska',
    'Nevada',
    'New Hampshire',
    'New Jersey',
    'New Mexico',
    'New York',
    'North Carolina',
    'North Dakota',
    'Ohio',
    'Oklahoma',
    'Oregon',
    'Pennsylvania',
    'Rhode Island',
    'South Carolina',
    'South Dakota',
    'Tennessee',
    'Texas',
    'Utah',
    'Vermont',
    'Virginia',
    'Washington',
    'West Virginia',
    'Wisconsin',
    'Wyoming',
  ]))