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
//Distribution 33.33%,runtime 116ms
var postorderTraversal = function(root) {
    var stack = [];
    var list = [];
    if(root === null) return list;
    stack.unshift(root);
    while(stack.length){
        var tree= stack.shift();  
        list.unshift(tree.val);
        if(tree.left !== null)
        stack.unshift(tree.left);
        if(tree.right !== null)
        stack.unshift(tree.right);
    }
    return list;
};
//Distribution 33.33%,runtime 116ms
var postorderTraversal = function(root) {
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
        if(tree.pass !== undefined){
            list.push(tree.val);
            tree.left = null;
        }else{
            tree.pass = true;
            stack.unshift(tree);
            if(tree.right !== null){
                tree = tree.right;
                stack.unshift(tree);
            }else{
                tree.left = null;
            }
        }
    }
    return list;
};
//Distribution 23.53%,runtime 124ms
var postorderTraversal = function(root) {
    var list = [];
    var preorder = function(root){
        if(root){
            preorder(root.left);
            preorder(root.right);
            list.push(root.val);
        }
    }
    preorder(root);
    return list;
};