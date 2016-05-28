/**
 * @param {number} n
 * @return {string}
 */
var countAndSay = function(n) {
    var list = "1";
    var node = "";
    for(var i = 1; i < n; i++){
        for(var j = 0, k = 0; j < list.length; j++){
            if(!list[j+1] || list[j] !== list[j+1]){
                node += list.substr(k,j-k+1).length + list.substr(k,1);
                k = j+1;
            }
        }
        list = node;
        node = "";
    }
    return list;
};