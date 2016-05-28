/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isIsomorphic = function(s, t) {
    var list1 = {};
    for(var i = 0, n = s.length; i < n; i++){
        if(list1[s[i]]){
            if(list1[s[i]] !== t[i]) return false;
        }else{
            list1[s[i]] = t[i];
        }
    }
    var list2 = {};
    for(var j = 0; j < n; j++){
        if(list2[t[j]]){
            if(list2[t[j]] !== s[j]) return false;
        }else{
            list2[t[j]] = s[j];
        }
    }
    return true;
};