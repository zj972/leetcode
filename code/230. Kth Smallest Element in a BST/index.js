/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} k
 * @return {number}
 */
//Distribution 7.69%,runtime 176ms
var kthSmallest = function(root, k) {
    var st = [];
    while (root !== null) {
        st.push(root);
        root = root.left;
    }
    while (k !== 0) {
        var n = st.pop();
        k--;
        if (k === 0) return n.val;
        var right = n.right;
        while (right !== null) {
            st.push(right);
            right = right.left;
        }
    }
    return -1;
};
//Distribution 5.77%,runtime 192ms
var kthSmallest = function(root, k) {
    var countNodes = function(n) {
        if (n === null) return 0;
        return 1 + countNodes(n.left) + countNodes(n.right);
    }
    var count = countNodes(root.left);
    if(k <= count){
        return kthSmallest(root.left, k);
    }else if(k > count + 1){
        return kthSmallest(root.right, k-1-count);
    }
    return root.val;
};
//Distribution 3.85%,runtime 216ms
var kthSmallest = function(root, k) {
    var number = 0;
    var helper = function(n){
        if (n.left !== null) helper(n.left);
        k--;
        if (k === 0) {
            number = n.val;
            return;
        }
        if (n.right !== null) helper(n.right);
    }
    helper(root);
    return number;
};