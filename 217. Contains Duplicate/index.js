/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    if(nums === null) return false;
    nums.sort(function(a,b){ return a-b });
    for(var i = 1,n = nums.length;i < n;i++){
        if(nums[i] === nums[i-1]) return true;
    }
    return false;
};