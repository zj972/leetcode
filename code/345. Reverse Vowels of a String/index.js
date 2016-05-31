/**
 * @param {string} s
 * @return {string}
 */
//Distribution 92.56%,runtime 152ms
var reverseVowels = function(s) {
    s = s.split("");
    var x = 0;
    var y = s.length;
    while(1){
        if(x >= y) return s.join("");
        if(s[x] === "a" || s[x] === "e" || s[x] === "i" || s[x] === "o" || s[x] === "u" || s[x] === "A" || s[x] === "E" || s[x] === "I" || s[x] === "O" || s[x] === "U"){
            if(s[y] === "a" || s[y] === "e" || s[y] === "i" || s[y] === "o" || s[y] === "u" || s[y] === "A" || s[y] === "E" || s[y] === "I" || s[y] === "O" || s[y] === "U"){
                var temp = s[x];
                s[x] = s[y];
                s[y] = temp;
                x++;
                y--;
            }else{
                y--;
            }
        }else{
            x++;
        }
    }
};