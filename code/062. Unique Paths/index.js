/**
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
//Distribution 89.58%,runtime 96ms
var uniquePaths = function(m, n) {
    var list = [];
    for(var x = 0; x < m; x++)
        list[x] = [];
    list[0][0] = 1;
    for(x = 0; x < m; x++){
        for(var y = 0; y < n; y++){
            if(x === 0 || y === 0) list[x][y] = 1;
            else list[x][y] = list[x - 1][y] + list[x][y - 1];
        }
    }
    return list[m - 1][n - 1];
};