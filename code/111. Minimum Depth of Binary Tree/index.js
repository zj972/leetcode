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
//Distribution 11.36%,runtime 144ms
var minDepth = function(root) {
    if (!root) return 0;
    if (root.left &&  root.right) return Math.min(minDepth(root.left), minDepth(root.right)) + 1;
    return Math.max(minDepth(root.left), minDepth(root.right)) + 1;
};