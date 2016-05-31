/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
//Distribution 53.75%,runtime 116ms
var maxDepth = function(root) {
    if(root == null){
        return 0;
    }else{
        return 1 + Math.max(maxDepth(root.left),maxDepth(root.right));
    }
};