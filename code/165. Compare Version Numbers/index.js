/**
 * @param {string} version1
 * @param {string} version2
 * @return {number}
 */
//Distribution 52.11%,runtime 104ms
var compareVersion = function(version1, version2) {
    version1 = version1.split(".");
    version2 = version2.split(".");
    for(var i = 0, n = Math.max(version1.length,version2.length); i < n; i++){
        if(!version1[i] && ! version2[i]) return 0;
        else if(!version1[i] && parseInt(version2[i])) return -1;
        else if(parseInt(version1[i]) && !version2[i]) return 1;
        else if(parseInt(version1[i]) > parseInt(version2[i])) return 1;
        else if(parseInt(version1[i]) < parseInt(version2[i])) return -1;
    }
    return 0;
};