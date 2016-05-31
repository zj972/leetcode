/**
 * @param {number} n - a positive integer
 * @return {number}
 */
//Distribution 11.89%,runtime 172ms
var hammingWeight = function(n) {
    n = n.toString(2);
    for(var i = 0, x = n.length,num = 0; i < x; i ++){
        if(n[i] === "1"){
            num++;
        }
    }
    return num;
};