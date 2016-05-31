/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
//Distribution 28.05%,runtime 120ms
var removeElement = function(nums, val) {
    var valList = 0;
    for(var i = 0, n = nums.length; i < n; i++){
        if(nums[i] === val){
            nums.splice(i,1);
            i--;
        }else{
            valList++;
        }
    }
    return valList;
};