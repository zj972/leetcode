/**
 * @param {number[]} nums
 * @return {number}
 */
//Distribution 31.91%,runtime 116ms
var singleNumber = function(nums) {
    var ans = 0;
    for (var i = 0; i < 32; i++){
        var cnt = 0, bit = 1 << i;
        for (var j = 0, n = nums.length; j < n; j++){
            if (nums[j] & bit) cnt++;
        }
        if (cnt % 3 !== 0)
            ans |= bit;
    }
    return ans;
    
};
//位运算，这里利用二进制模拟a进制的方法。对于此题，a为3.
//对于除出现一次之外的所有的整数，其二进制表示中每一位1出现的次数是3的整数倍，将所有这些1清零，剩下的就是最终的数。
//用ones记录到当前计算的变量为止，二进制1出现“1次”（mod 3 之后的 1）的数位。
//用twos记录到当前计算的变量为止，二进制1出现“2次”（mod 3 之后的 2）的数位。
//当ones和twos中的某一位同时为1时表示二进制1出现3次，此时需要清零。即用二进制模拟三进制计算。
//最终ones记录的是最终结果。
//Distribution 29.79%,runtime 120ms
var singleNumber = function(nums) {
    var ones = 0, twos = 0, xthrees = 0;
    for(var i = 0, n = nums.length; i < n; i++) {
        twos |= (ones & nums[i]);
        ones ^= nums[i];
        xthrees = ~(ones & twos);
        ones &= xthrees;
        twos &= xthrees;
    }
    return ones;
};
//Distribution 17.02%,runtime 128ms
var singleNumber = function(nums) {
    var map = {};
    for(var i = 0, n = nums.length; i < n; i++){
        if(map[nums[i]] === undefined) map[nums[i]] = 1;
        else map[nums[i]]++;
    }
    for(i = 0; i < n; i++){
        if(map[nums[i]] === 1) return nums[i];
    }
};