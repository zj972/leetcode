/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    for(var i = 0; i < nums.length; i++){
        if(i+1 === nums.length){
            break;
        }else{
            if(nums[i] === nums[i+1]){
                nums.splice(i+1, 1);
                i--;
            }
        }
    }
};