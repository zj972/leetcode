/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
//Distribution 52.68%,runtime 112ms
var invertTree = function(root) {
    if(root == null) return [];
    var ec;
    ec = root.left;
    root.left = root.right;
    root.right = ec;
    invertTree(root.left);
    invertTree(root.right);
    return root;
};