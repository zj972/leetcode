/**
 * @param {number[]} nums
 * @return {number}
 */
//Distribution 63.74%,runtime 104ms
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
//Distribution 41.52%,runtime 116ms
var singleNumber = function(nums) {
    var list = 0;
    for(var i = 0, n = nums.length; i < n; i++){
        list ^= nums[i];
    }
    return list;
};