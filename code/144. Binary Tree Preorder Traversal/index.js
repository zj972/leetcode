/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
//Distribution 82.14%,runtime 108ms
var preorderTraversal = function(root) {
    var list = [];
    var preorder = function(root){
        if(root){
            list.push(root.val);
            preorder(root.left);
            preorder(root.right);
        }
    }
    preorder(root);
    return list;
};