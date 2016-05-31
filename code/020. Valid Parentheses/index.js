/**
 * @param {string} s
 * @return {boolean}
 */
//Distribution 71.83%,runtime 96ms
var isValid = function(s) {
    var list = [];
    for(var i = 0, n = s.length; i < n; i++){
        switch(s[i]){
            case "(" :
                list.push(s[i])
                break;
            case ")" :
                if(list[list.length -1] === "("){
                    list.length--;
                }else{
                    return false;
                }
                break;
            case "{" :
                list.push(s[i])
                break;
            case "}" :
                if(list[list.length -1] === "{"){
                    list.length--;
                }else{
                    return false;
                }
                break;
            case "[" :
                list.push(s[i])
                break;
            case "]" :
                if(list[list.length -1] === "["){
                    list.length--;
                }else{
                    return false;
                }
                break;
            default  : break;
        }
    }
    return !list.length;
};