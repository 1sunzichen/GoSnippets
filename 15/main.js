//最长前缀
function long(strs){
    let str=strs[0]
    while (!strs.every((item)=>item.startsWith(str))) {
            str=str.slice(0,str.length-1)                          
    }
    return str
}
let str=["fi","foo","foi"]
console.log(long(str));