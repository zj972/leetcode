/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
//Distribution 75.79%,runtime 128ms
var addBinary = function(a, b) {
    if(a.length > b.length){
        var temp = a;
        a = b;
        b = temp;
    }
    a = a.split("");
    b = b.split("");
    for(var i = 0, n = b.length - a.length; i < n; i++){
        a.unshift("0");
    }
    for(var j = a.length - 1, go = 0; j >= 0; j--){
        go += parseInt(a[j]) + parseInt(b[j]);
        switch(go){
            case 0 :{a[j] = "0";go = 0;break;}
            case 1 :{a[j] = "1";go = 0;break;}
            case 2 :{a[j] = "0";go = 1;break;}
            case 3 :{a[j] = "1";go = 1;break;}
        }
        if(!a[j-1] && go){
            a.unshift("1");
        }
    }
    return a.join("");
};