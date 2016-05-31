/**
 * @param {number} num
 * @return {number}
 */
//Distribution 37.87%,runtime 184ms
var addDigits = function(num) {
    return (num-1)%9+1;
};