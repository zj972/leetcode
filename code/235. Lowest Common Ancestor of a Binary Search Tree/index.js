/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
//Distribution 62.63%,runtime 160ms
var lowestCommonAncestor = function(root, p, q) {
    while((root.val - p.val) * (root.val - q.val) > 0){
        root = p.val < root.val ? root.left : root.right;
    }
    return root;
};