/**
 * @param {number[]} nums
 * @return {number}
 */
 //位运算 a^a=0,a^0=a
//Distribution 55.26%,runtime 125ms
var missingNumber = function(nums) {
    for (var i = 1, res = 0, n = nums.length; i <= n; i++)   
            res = res ^ i ^ nums[i-1];
    return res;    
};
//Distribution 23.68%,runtime 140ms
var missingNumber = function(nums) {
    var sum = 0;    
    for(var i in nums)  
        sum += nums[i];    
    var n = nums.length;    
    return (n * (n + 1))/ 2 - sum;    
};