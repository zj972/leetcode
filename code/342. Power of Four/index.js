/**
 * @param {number} num
 * @return {boolean}
 */
//Distribution 2.65%,runtime 212ms
var isPowerOfFour = function(num) {
    if(num === 1){
        return true;
    }else if(num < 4){
        return false;
    }else{
        num = num/4;
        if(num === parseInt(num)){
            return isPowerOfFour(num);
        }else{
            return false;
        }
    }
};