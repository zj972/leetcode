/**
 * @param {number} num
 * @return {string}
 */
//Distribution 63.37%,runtime 332ms
var intToRoman = function(num) {
    var roman = "";
    for(var i = 0; i < 4; i++){
        var w = Math.pow(10,3-i);
        var n = parseInt(num / w);
        num = num - n * w;
        var x = ["M","C","X","I"];
        var y = ["","D","L","V"];
        if(n === 9){
            roman += x[i] + x[i - 1];
        }else if(n > 4){
            roman += y[i];
            for(var j = 0; j < n - 5; j++){
                roman += x[i];
            }
        }else if(n === 4){
            roman += x[i] + y[i];
        }else{
            for(var k = 0; k < n; k++){
                roman += x[i];
            }
        }
    }
    return roman;
};
//Distribution 42.57%,runtime 364ms
var intToRoman = function(num) {
    var M = ["", "M", "MM", "MMM"];
    var C = ["", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"];
    var X = ["", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"];
    var I = ["", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"];
    return M[parseInt(num/1000)] + C[parseInt((num%1000)/100)] + X[parseInt((num%100)/10)] + I[num%10];
};