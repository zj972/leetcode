/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
//Distribution 87.50%,runtime 116ms
var moveZeroes = function(nums) {
    var zeroCount = 0;
    for(var i = 0 ; i< nums.length; i++) {
        if (nums[i] === 0) {
            zeroCount++;
        } else if (zeroCount > 0 ){
            nums[i-zeroCount] = nums[i];
            nums[i] = 0;
        }
    }
};