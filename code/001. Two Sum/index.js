/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    var num = {};
    for(var i = 0, n = nums.length; i < n; i++){
        var end = target - nums[i];
        if(num[end] !== undefined){
            if(i > num[end]){
                return [num[end],i]
            }else{
                return [i,num[end]]
            }
        }else{
            num[nums[i]] = i;
        }
    }
};