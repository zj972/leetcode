/**
 * @param {character[][]} board
 * @return {boolean}
 */
//Distribution 58.11%,runtime 172ms
var isValidSudoku = function(board) {
    for(var i = 0; i < 9; i++){
        var col = [];
        for(var s = 0; s < 9; s++){
            col.push(board[i][s]);
        }
        col.sort();
        for(var j = 0; j < 9; j++){
            if(col[j] !== "." && col[j] === col[j+1]){
                return false;
            }
        }
    }
    for(var k = 0; k < 9; k++){
        var row = [];
        for(var l = 0; l < 9; l++){
            row.push(board[l][k]);
        }
        row.sort();
        for(var m = 0; m < 9; m++){
            if(row[m] !== "." && row[m] === row[m+1]){
                return false;
            }
        }
    }
    for(var a = 0; a < 3; a++){
        for(var b = 0; b < 3; b++){
            var z = [];
            for(var c = 0; c < 3; c++){
                for(var d = 0; d < 3; d++){
                    z.push(board[a*3 + c][b*3 + d]);
                }
            }
            z.sort();
            for(var e = 0; e < 9; e++){
                if(z[e] !== "." && z[e] === z[e+1]){
                    return false;
                }
            }
        }
    }
    return true;
};