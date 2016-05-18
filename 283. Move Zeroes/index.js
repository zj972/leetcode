/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function(nums) {
    var time = 0;
    for(var i = 0;i < nums.length;i++){
        if(nums[i] == 0){
            nums.splice(i,1);
            var zero = 0;
            nums.push(zero);
            i--;
            time++;
            if(time == nums.length) break;
        }else{
            time = 0;
        }
    }
};