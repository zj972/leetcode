/**
 * @param {string} s
 * @return {string}
 */
var reverseString = function(s) {
    s = s.split("").reverse().join("");
    return s;
};