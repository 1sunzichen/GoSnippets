//最长公共前缀
//https://leetcode.cn/problems/longest-common-prefix/
function test(strs){
    if(strs.length==0){
        return ""
    }
    let ans=strs[0]
    for(let i=1;i<strs.length;i++){
        let j=0
        for(;j<ans.length&&j<strs[i].length;j++){
            if(ans[j]!=strs[i][j]) break
        }
        ans=ans.substr(0,j)
        if(ans===""){
            return ""
        }
    }
    return ans
}
console.log(test(["afd","afdfd","afdff"]))