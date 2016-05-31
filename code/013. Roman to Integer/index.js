/**
 * @param {string} s
 * @return {number}
 */
//Distribution 34.07%,runtime 364ms
var romanToInt = function(s) {
    var x = 0;
    for(var i = 0, num = [], n = s.length; i < n; i++){
        switch(s[i]){
            case "I": num[i] = 1;break;
            case "V": num[i] = 5;break;
            case "X": num[i] = 10;break;
            case "L": num[i] = 50;break;
            case "C": num[i] = 100;break;
            case "D": num[i] = 500;break;
            case "M": num[i] = 1000;break;
        }
    }
    for(var j = 0, rules = 0, m = num.length; j < m; j++){
        if(num[j+1] === null){
            x += num[j];
            continue;
        }
        if(num[j] > num[j+1]){
            if(rules){
                x += (rules + 1) * num[j];
                rules = 0;
            }else{
                x += num[j];   
            }
        }else if(num[j] === num[j+1]){
            rules++;
        }else if(j+1 === m){
            if(rules){
                x += (rules + 1) * num[j];
                rules = 0;
            }else{
                x += num[j];
            }
        }else{
            if(rules){
                x -= (rules + 1) * num[j];
                rules = 0;
            }else{
                x -= num[j];
            }
        }
    }
    return x;
};