/**
 * @param {number} n
 * @return {number}
 */
//Distribution 38.18%,runtime 96ms
var numTrees = function(n) {
    var f = [1,1];
    for(var a = 2; a <= n; f[a++] = 0); 
    for (var i = 2; i <= n; i++){  
        for (var k = 1; k <= i; k++)
            f[i] = f[i] + f[k - 1] * f[i - k];
    }  
    return f[n];
};