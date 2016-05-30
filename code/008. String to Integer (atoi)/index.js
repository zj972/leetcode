/**
 * @param {string} str
 * @return {number}
 */
//runtime 56.25% 180ms
var myAtoi = function(str){
    var sign = 1, base = 0, i = 0;
    while (str[i] === " "){
        i++;
    }
    if (str[i] === "-" || str[i] === "+"){
        sign = 1 - 2 * (str[i++] === "-"); 
    }
    while (str[i] >= "0" && str[i] <= "9"){
        var max = Math.floor(2147483647 / 10);
        if (base > max || (base === max && str[i] - "0" > 7)) {
            if (sign === 1) return 2147483647;
            else return -2147483648;
        }
        base  = 10 * base + (str[i++] - '0');
    }
    return base * sign;
};
//runtime 14.01% 202ms
// var myAtoi = function(str) {
//     str = str.split("");
//     for(var i = 0, n = str.length, key = 0, num = ""; i < n; i++){
//         if(str[i] === " " && key === 0) continue;
//         if(key === 1){
//             if(str[i].match(/[0-9]/) !== null){
//                 num += str[i];
//             }else if(str[i].match(/[-\+]/) !== null){
//                 return 0;
//             }else{
//                 num = parseInt(num);
//                 if(num !== num) return 0;
//                 return (num>0)?(num<2147483647?num:2147483647):(num>-2147483648?num:-2147483648);
//             }
//         }else if(key === 0){
//             if(str[i].match(/[0-9]/) !== null){
//                 num = str[i];
//                 key = 1;
//             }else if(str[i].match(/-/) !== null){
//                 num = "-";
//                 key = 1;
//             }else if(str[i].match(/\+/) !== null){
//                 num = "";
//                 key = 1;
//             }else return 0;
//         }
//     }
//     num = parseInt(num)?parseInt(num):0;
//     return (num>0)?(num<2147483647?num:2147483647):(num>-2147483648?num:-2147483648);
// };