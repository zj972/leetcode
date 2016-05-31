/**
 * @param {number} n
 * @return {boolean}
 */
//Distribution 11.98%,runtime 192ms
var isPowerOfTwo = function(n) {
    if(n === 1){
        return true;
    }else if(n < 2){
        return false;
    }else{
        n = n/2;
        if(n === parseInt(n)){
            return isPowerOfTwo(n);
        }else{
            return false;
        }
    }
};
//log函数
//Distribution 11.98%,runtime 192ms
var isPowerOfThree = function(n) {
    var logAns = Math.log(n) / Math.log(2);
    return (Math.abs(logAns - Math.round(logAns)) < 10e-15) ? true : false;
};