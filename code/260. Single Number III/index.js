/**
 * @param {number[]} nums
 * @return {number[]}
 */
//Distribution 69.23%,runtime 116ms
var singleNumber = function(nums) {  
    var res = [0,0];  
    var result = nums[0];  
    for(var i=1;i<nums.length;i++){  
        result = result^nums[i];  
    } 
    var n = result & (~(result-1));  
    for(var j=0;j<nums.length;j++){  
        if((n & nums[j])!==0){  
            res[0] = res[0] ^ nums[j];  
        }else {  
            res[1] = res[1] ^ nums[j];  
        }  
    }  
    return res;  
};
//Distribution 34.07%,runtime 152ms
var singleNumber = function(nums) {
    var list = {};
    for(var i = 0, n = nums.length; i < n; i++){
        if(list[nums[i]] === undefined){
            list[nums[i]] = i;
        }else{
            delete list[nums[i]];
        }
    }
    var num = [];
    for(var x in list) num.push(parseInt(x));
    if(list[num[0]] < list[num[1]]) return num;
    else return num.reverse();
};