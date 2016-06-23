/**
 * @param {number[]} prices
 * @return {number}
 */
//Distribution 20.00%,runtime 144ms
var maxProfit = function(prices) {
    var num = prices.length;
    if(num <= 1) return 0;
    for(var i = 0, s0 = [], s1 = [], s2 = []; i < num; i++){
        s0[i] = 0;
        s1[i] = 0;
        s2[i] = 0;
    }
    s1[0] = -prices[0];
    s2[0] = Number.MIN_VALUE;
    for(i = 1; i < num; i++){
        s0[i] = Math.max(s0[i - 1], s2[i - 1]);
        s1[i] = Math.max(s1[i - 1], s0[i - 1] - prices[i]);
        s2[i] = s1[i - 1] + prices[i];
    }
    return Math.max(s0[num - 1], s2[num - 1]);
};