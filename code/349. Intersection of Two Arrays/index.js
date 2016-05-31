/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
//Distribution 81.63%,runtime 112ms
var intersection = function(nums1, nums2) {
    var sam = [];
    var b = 0;
    for(var i = 0;i < nums2.length;i++){
        for(var j = 0;j < nums1.length;j++){
            if(nums2[i] === nums1[j]){
                b = 1;
            }
        }
        if(b === 1){
            for(var x = 0,c = 0;x < sam.length;x++){
                if(nums2[i] === sam[x]){
                    c = 1;  
                }
            }
            if(!c){
                sam.push(nums2[i]);
            }
            b = 0;
        }
    }
    return sam;
};