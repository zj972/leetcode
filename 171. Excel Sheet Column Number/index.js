/**
 * @param {string} s
 * @return {number}
 */
var titleToNumber = function(s) {
    for(var i = 0,ret = 0; i < s.length; i ++){
        ret = ret*26 + (s.charCodeAt(i)-"A".charCodeAt(0)+1);
    }
    return ret;
};