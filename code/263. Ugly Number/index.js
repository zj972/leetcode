/**
 * @param {number} num
 * @return {boolean}
 */
var isUgly = function(num) {
    if(num === 0) return false;
    if(num === 1) return true;
    if(parseInt(num/2) === (num/2)){
        return isUgly(num/2);
    }else if(parseInt(num/3) === (num/3)){
        return isUgly(num/3);
    }else if(parseInt(num/5) === (num/5)){
        return isUgly(num/5);
    }else{
        return false;
    }
};