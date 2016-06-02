/**
 * @param {string[]} words
 * @return {number}
 */
//Distribution 39.13%,runtime 276ms
var maxProduct = function(words) {
    var list = [];
    for(var i = 0, n = words.length; i < n; i++){
        var word = words[i].split("");
        for(var j = 0, m = word.length; j < m; j++){
            list[i] |= 1 << (word[j].charCodeAt() - 97); 
        }
    }
    var max = 0;
    var num = 0;
    for(var k = 0; k < n; k++){
        for(var l = k + 1; l < n; l++){
            if((list[k] & list[l]) === 0){
                num = words[k].length * words[l].length;
                max = max > num ? max : num;
            }
        }
    }
    return max;
};