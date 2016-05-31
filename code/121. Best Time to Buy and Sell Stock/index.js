/**
 * @param {number[]} prices
 * @return {number}
 */
//Distribution 14.19%,runtime 140ms
var maxProfit = function(prices) {
    var n = prices.length;
    var min = Number.MAX_VALUE;
    var money = 0;
    for(var i = 0; i < n; i++){
        if(prices[i] < min){
            min = prices[i];
        }else{
            if((prices[i] - min) > money){
                money = prices[i] - min;
            }
        }
    }
    return money;
};