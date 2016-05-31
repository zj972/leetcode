/**
 * @param {string} pattern
 * @param {string} str
 * @return {boolean}
 */
//Distribution 55.75%,runtime 97ms
var wordPattern = function(pattern, str) {
    var list1 = {};
    str = str.split(" ");
    if(pattern.length !== str.length) return false;
    for(var i = 0, n = pattern.length; i < n; i++){
        if(list1[pattern[i]]){
            if(list1[pattern[i]] !== str[i]) return false;
        }else{
            list1[pattern[i]] = str[i];
        }
    }
    var list2 = {};
    for(var j = 0; j < n; j++){
        if(list2[str[j]]){
            if(list2[str[j]] !== pattern[j]) return false;
        }else{
            list2[str[j]] = pattern[j];
        }
    }
    return true;
};