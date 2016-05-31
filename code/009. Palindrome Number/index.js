/**
 * @param {number} x
 * @return {boolean}
 */
//Distribution 50.00%,runtime 756ms
var isPalindrome = function(x) {
    if(x < 0) return false;
    else if(x === 0) return true;
    xString = x.toString();
    var y = parseInt(xString.split("").reverse().join(""));
    if(x === y) return true;
    else return false;
};