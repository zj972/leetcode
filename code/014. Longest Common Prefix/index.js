/**
 * @param {string[]} strs
 * @return {string}
 */
//Distribution 55.79%,runtime 108ms
var longestCommonPrefix = function(strs) {
    if(!strs.length) return "";
    var str = strs[0];
    for(var i = 1; i < strs.length; i++){
        for(var j = 0,n = Math.max(strs[i].length,str.length); j < n; j++){
            if(!strs[i][j]){
                str = str.substr(0,j);
            }else if(!str[j]){
                continue;
            }else if(str[j] !== strs[i][j]){
                str = str.substr(0,j);
            }
        }
    }
    return str;
};