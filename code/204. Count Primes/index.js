/**
 * @param {number} n
 * @return {number}
 */
 /**
 * @param {number} n
 * @return {number}
 */
//Distribution 63.89%,runtime 376ms
var countPrimes = function(n) {
    if(n <= 2){
        return 0;
    }
    var mem = [];
    for(var i = 2; i < n; i++){
        mem[i] = true;
    }
    sq = parseInt(Math.sqrt(n - 1));
    for(i = 2; i <= sq; i++){
        if(mem[i]){
            for(var j = i + i; j < mem.length; j += i){
                mem[j] = false;
            }
        }
    }
    var count = 0;
    for(i = 2; i < mem.length; i++){
        if(mem[i]){
            count++;    
        }
    }
    return count;
};
//Distribution 36.11%,runtime 664ms
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