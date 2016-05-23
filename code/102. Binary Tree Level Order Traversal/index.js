/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
var levelOrder = function(root) {
    var num = 0;
    var rootArray = [];
    function level(root){
        if(root){
            if(rootArray[num] === undefined) rootArray[num] = [];
            rootArray[num].push(root.val);
            num++;
            level(root.left);
            level(root.right);
            num--;
        }
    }
    level(root);
    return rootArray;
};