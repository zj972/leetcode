/**
 * @param {number} n
 * @return {number}
 */
//Distribution 22.92%,runtime 108ms
var integerBreak = function(n) {
    if(n <= 3) return n-1;
    var mod = n % 3;
    if(mod === 0) return Math.pow(3,parseInt(n/3));
    else if(mod === 1) return 4 * Math.pow(3, parseInt((n-4)/3));
    else return 2 * Math.pow(3, parseInt(n/3));
};