/**
 * @param {number[]} nums
 * @return {number}
 */
//Distribution 71.79%,runtime 96ms
var rob = function(nums) {
    if (nums === null || nums.length === 0) {
        return 0;
    }
    if (nums.length > 2) {
        nums[2] += nums[0];
    }
    for (var i = 3; i < nums.length; i++) {
        nums[i] += Math.max(nums[i - 2], nums[i - 3]);
    }
    if (nums.length == 1) {
        return nums[0];
    }else if (nums.length == 2) {
        return Math.max(nums[0], nums[1]);
    }else {
        return Math.max(nums[i - 1], nums[i - 2]);
    }
};