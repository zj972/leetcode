/**
 * @param {number} numRows
 * @return {number[][]}
 */
var generate = function(numRows) {
    if(!numRows) return [];
    var numArray = [[1]];
    for(var i = 1; i < numRows; i++){
        numArray[i] = [1];
        for(var j = 1; j < i+1; j++){
            if(j === i){
                numArray[i][j] = 1;
            }else{
                numArray[i][j] = numArray[i-1][j-1] + numArray[i-1][j];
            }
        }
    }
    return numArray;
};