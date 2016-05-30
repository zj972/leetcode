/**
 * @param {number[]} nums
 * @return {number}
 */
//runtime 63.74% 104ms
var singleNumber = function(nums) {
    var list = {};
    for(var i = 0, n = nums.length; i < n; i++){
        if(list[nums[i]] === undefined){
            list[nums[i]] = nums[i];
        }else{
            delete list[nums[i]];
        }
    }
    for(var x in list) return list[x];
};
//runtime 41.52% 116ms
var singleNumber = function(nums) {
    var list = 0;
    for(var i = 0, n = nums.length; i < n; i++){
        list ^= nums[i];
    }
    return list;
};