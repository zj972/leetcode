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
//递归
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
//迭代
//Distribution 47.62%,runtime 119ms
var preorderTraversal = function(root) {
    var list = [];
    var stack = [];
    if(root === null) return list;
    stack.unshift(root);
    while(stack.length > 0){
        tree = stack.shift();
        list.push(tree.val);
        if(tree.right !== null) stack.unshift(tree.right);
        if(tree.left !== null) stack.unshift(tree.left);
    }
    return list;
};