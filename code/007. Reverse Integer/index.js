/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
    var max = Math.floor(2147483647/10), result = 0;
    var sign = x > 0 ? 1 : -1;
    x *= sign;
    while(x > 0){
        if(result > max || (result === max && x > 7)) return 0;
        else{
            result = result * 10 + x % 10;
            x = Math.floor(x/10);
        }
    }
    return sign * result;
};