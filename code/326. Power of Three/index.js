/**
 * @param {number} n
 * @return {boolean}
 */
 //递归
//Distribution 26.01%,runtime 1252ms
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
//Distribution 23.77%,runtime 1260ms
var isPowerOfThree = function(n) {
    var logAns = Math.log(n) / Math.log(3);
    return (Math.abs(logAns - Math.round(logAns)) < 10e-15) ? true : false;
};