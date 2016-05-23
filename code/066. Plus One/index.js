/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function(digits) {
    if(!digits.length) return [1];
    var num = 1;
    var n = digits.length-1;
    while(num){
        if(n === -1){
            digits.unshift(1);
            num--;
        }else if(digits[n] === 9){
            digits[n] = 0;
            n--;
        }else{
            digits[n]++;
            num--;
        }
    }
    return digits;
};