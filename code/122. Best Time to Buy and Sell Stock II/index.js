/**
 * @param {number[]} prices
 * @return {number}
 */
//runtime 26.67% 132ms
var maxProfit = function(prices) {
    var n = prices.length;
    var max = prices[0];
    var min = prices[0];
    var money = 0;
    for(var i = 0; i < n; i++){
        if(prices[i] >= max){
            max = prices[i];
        }else{
            money += max - min;
            min = prices[i];
            max = prices[i];
        }
    }
    if(max !== min) return (money + max - min);
    return money;
};
//runtime 11.67% 148ms
// var maxProfit = function(prices) {
//     var n = prices.length;
//     var money = 0;
//     for(var i = 0; i < n; i++){
//         if (prices[i+1]>prices[i]) money += prices[i+1]-prices[i];
//     }
//     return money;
// };