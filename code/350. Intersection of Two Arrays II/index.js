/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersect = function(nums1, nums2) {
    var sam = [];
    var b = 0;
    for(var i = 0;i < nums2.length;i++){
        for(var j = 0;j < nums1.length;j++){
            if(nums2[i] === nums1[j]){
                sam.push(nums2[i]);
                nums1[j] = null;
                break;
            }
        }
    }
    return sam;
};