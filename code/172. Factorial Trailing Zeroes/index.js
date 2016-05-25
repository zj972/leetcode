/**
 * @param {number} n
 * @return {number}
 */
var trailingZeroes = function(n) {
    if(n < 5) return 0;
    n = parseInt(n/5);
    var num = n;
    while(n >= 5){
        num += parseInt(n/5);
        n = parseInt(n/5);
    }
    return num;
};