/**
 * @param {number[]} nums
 * @param {number} k
 * @return {void} Do not return anything, modify nums in-place instead.
 */
//Distribution 26.51%,runtime 224ms
var rotate = function(nums, k) {
    if(k){
        for(var i = 0, n = nums.length - 1; i < k; i++){
            var num = nums[n];
            nums.pop();
            nums.unshift(num);
        }
    }
};