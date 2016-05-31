/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
//Distribution 80.54%,runtime 100ms
var twoSum = function(nums, target) {
    var num = {};
    for(var i = 0, n = nums.length; i < n; i++){
        var end = target - nums[i];
        if(num[end] !== undefined){
            if(i > num[end]){
                return [num[end],i];
            }else{
                return [i,num[end]];
            }
        }else{
            num[nums[i]] = i;
        }
    }
};
//Distribution 51.73%,runtime 168ms
var twoSum = function(nums, target) {
    var num = [];
    for(var i = 0, n = nums.length; i < n; i++){
        var end = target - nums[i];
        for(var j = 0,m = num.length; j < m; j++){
            if(end === num[j]){
                if(i > j){
                    return [j,i];
                }else{
                    return [i,j];
                }
            }
        }
        num[i] = nums[i];
    }
};