/**
 * @param {number} n
 * @return {string}
 */
//Distribution 69.53%,runtime 88ms
var convertToTitle = function(n) {
    var list = "";
    var go = 0;
    while(n > 26){
        go = n % 26;
        if(go === 0){
            go = 26;
            n = n/26 - 1;
        }
        else n = (n - go)/26;
        list = String.fromCharCode(go+64) + list;
    }
    list = String.fromCharCode(n+64) + list;
    return list;
};
//Distribution 22.58%,runtime 100ms
var convertToTitle = function(n) {
    if (n === 0) return "";
    return convertToTitle(parseInt((n - 1) / 26)) + String.fromCharCode((n - 1) % 26 + 65);
};
