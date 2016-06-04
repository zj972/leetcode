/**
 * @param {number} num
 * @return {number[]}
 */
//Distribution 86.11%,runtime 284ms
var countBits = function(num) {
    var res = [0];
    for(var i=1;i<=num;i++){
        res[i] = res[i >> 1] + (i & 1);
    }
    return res;
};
//Distribution 61.81%,runtime 300ms
var countBits = function(num) {
    var res = [];
    var cnt = 1;
    var flag = 1;
    res.push(0);
    while(cnt <= num){
        var start = cnt;
        for (var tmp = flag; tmp > 0 && cnt <= num; --tmp){
            res.push(1 + res[start-tmp]);
            cnt++;
        }
        flag *= 2;
    }
    return res;
};