/**
 * @param {number} n
 * @return {number}
 */
//Distribution 99.21%,runtime 84ms
var countNumbersWithUniqueDigits = function(n) {
    if (n < 0) return 0;  
    if (n === 0) return 1;  
    if (n == 1) return 10;  
    var counts = 10;  
    var m = 9;  
    for(var i=1, j=9; i<n && j>0; i++, j--){  
        m *= j;  
        counts += m;  
    }  
    return counts;  
};