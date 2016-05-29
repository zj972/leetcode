/**
 * @param {number} n
 * @return {number}
 */
var countPrimes = function(n) {
    if(!n) return n;
    if(n === 1) return 0;
    var num = [];
    var m = 0;
    for(var i = 2; i < n; i++){
        if(!num[i]) m++;
        for(var j = i; j < n;){
            num[j] = 1;
            j += i;
        }
    }
    return m;
};