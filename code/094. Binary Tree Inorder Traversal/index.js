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
//迭代
//Distribution 89.04%,runtime 100ms
var inorderTraversal = function(root) {
    var list = [];
    if(root === null) return list;
    var stack = [];
    stack.unshift(root);
    var tree = stack[0];
    while(stack.length > 0){
        while(tree.left !== null){
            tree = tree.left;
            stack.unshift(tree);
        }
        tree = stack.shift();
        list.push(tree.val);
        if(tree.right !== null){
            tree = tree.right;
            stack.unshift(tree);
        }else{
            tree.left = null;
        }
    }
    return list;
};
//递归
//Distribution 47.95%,runtime 112ms
var inorderTraversal = function(root) {
    var list = [];
    var preorder = function(root){
        if(root){
            preorder(root.left);
            list.push(root.val);
            preorder(root.right);
        }
    }
    preorder(root);
    return list;
};