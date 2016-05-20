/**
 * @param {number} n
 * @return {boolean}
 */
 //递归
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
//log函数
// var isPowerOfThree = function(n) {
//     var logAns = Math.log(n) / Math.log(3);
//     return (Math.abs(logAns - Math.round(logAns)) < 10e-15) ? true : false;
// };