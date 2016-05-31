/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
//Distribution 44.96%,runtime 168ms
var topKFrequent = function(nums, k) {
    var map = {};
    for(var i = 0, n = nums.length; i < n; i++){
        if(map[nums[i]] === undefined){
            map[nums[i]] = 1;
        }else{
            map[nums[i]]++;
        }
    }
    var list = [];
    for(var j in map){
        list.push([parseInt(j),map[j]]);
    }
    list = list.sort(function(a,b){ return b[1]-a[1] });
    for(var l = 0,number = []; l < k; l++){
        number.push(list[l][0]);
    }
    return number;
};