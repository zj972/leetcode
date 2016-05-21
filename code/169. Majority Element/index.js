/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function(nums) {
    if(nums.length < 3) return nums[0];
    nums = nums.sort(function(a,b){ return a-b; });
    for(var i = 0,num = 0,element = nums[0];i < nums.length;i++){
        if(element === nums[i]){
            num++;
        }else{
            if(num > nums.length/2) return element;
            num = 1;
            element = nums[i];
        }
    }
    return element;
};