/**
 * @param {number} n
 * @return {boolean}
 */
var isPowerOfThree = function(n) {
    if(n === 1){
        return true;
    }else if(n < 3){
        return false;
    }else{
        n = n/3;
        if(n === parseInt(n)){
            return isPowerOfThree(n);
        }else{
            return false;
        }
    }
};