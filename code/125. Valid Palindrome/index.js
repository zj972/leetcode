/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function(s) {
    s = s.toLowerCase().split("");
    var m = s.length -1;
    var reg = /[a-zA-Z0-9]/;
    for(var i = 0, n = s.length; i < n; i++){
        while(!reg.test(s[i])){
            i++;
        }
        while(!reg.test(s[m])){
            m--;
        }
        if(m <= i) break;
        if(s[i] !== s[m]) return false;
        else m--;
    }
    return true;
};