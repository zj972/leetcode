/**
 * @param {number} rowIndex
 * @return {number[]}
 */
var getRow = function(rowIndex) {
    var numArray = [[1]];
    for(var i = 1; i < rowIndex+1; i++){
        numArray[i] = [1];
        for(var j = 1; j < i+1; j++){
            if(j === i){
                numArray[i][j] = 1;
            }else{
                numArray[i][j] = numArray[i-1][j-1] + numArray[i-1][j];
            }
        }
    }
    return numArray[rowIndex];
};