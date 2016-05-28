/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    s = s.trim().split(" ");
    if(s.length === 0) return s;
    else return s[s.length - 1].length;
};