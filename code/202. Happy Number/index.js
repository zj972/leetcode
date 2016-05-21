/**
 * @param {number} n
 * @return {boolean}
 */
var isHappy = function(n) {
    var list = [];
    list.push(n);
    while(1){
        var num = 0;
        n = n.toString();
        for(var i = 0, m = n.length; i < m; i++){
            num += n[i] * n[i];
        }
        n = num;
        if(n === 1) return true;
        for(var j = 0, x = list.length; j < x; j++){
            if(n === list[j]) return false;
        }
        list.push(n);
    }
};