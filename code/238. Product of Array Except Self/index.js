/**
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function(nums) {
    if(nums.length < 3) return nums.reverse();
    var p = [1], q = [];
    q[nums.length-1] = 1;
    for(var i = 1, n = nums.length; i < n; i++){
        p[i] = p[i-1] * nums[i-1];
        q[n-i-1] = q[n-i] * nums[n-i];
    }
    var number = [];
    number[0] = q[0];
    number[n-1] = p[n-1];
    for(var j = 1; j < n-1; j++){
        number[j] = p[j] * q[j];
    }
    return number;
};