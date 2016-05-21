/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isBalanced = function(root) {
    if(root === null) return true;
    function depth(root){
        if(root === null) return 0;
        return Math.max(depth(root.left),depth(root.right))+1
    }
    var left = depth(root.left);
    var right = depth(root.right);
    return Math.abs(left - right) <= 1 && isBalanced(root.left) && isBalanced(root.right);
};