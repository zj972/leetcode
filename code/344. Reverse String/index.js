/**
 * @param {string} s
 * @return {string}
 */
//Distribution 54.36%,runtime 156ms
var reverseString = function(s) {
    s = s.split("").reverse().join("");
    return s;
};