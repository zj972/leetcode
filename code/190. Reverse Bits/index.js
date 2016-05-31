/**
 * @param {number} n - a positive integer
 * @return {number} - a positive integer
 */
//Distribution 29.55%,runtime 164ms
var reverseBits = function(n) {
    var two = parseInt(n).toString(2).split("");
    for(var i = 0, m = two.length; i < 32 - m; i++){
    	two.unshift("0");
    }
    two = two.reverse();
    var ten = parseInt(two.join(""),2);
    return ten;
};