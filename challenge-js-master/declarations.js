const escapeStr = "`'\"\\/'"
const arr = Array(4,'2')
const obj = {str : "oki dokie",
            num : 8,
            bool : false,
            undef : undefined,
}
const nested = {arr : Array(4,undefined,'2'),
obj : {str : "TG",
num : 45,
bool : true,}}
Object.freeze(nested)
Object.freeze(nested.obj)
Object.freeze(nested.arr)