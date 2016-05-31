/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {string[]}
 */
//Distribution 67.90%,runtime 132ms
var binaryTreePaths = function(root) {
    function searchBT(root, path, answer) {
        if (root.left === null && root.right === null) answer.push(path + root.val);
        if (root.left !== null) searchBT(root.left, path + root.val + "->", answer);
        if (root.right !== null) searchBT(root.right, path + root.val + "->", answer);
    }
    var answer = [];
    if (root !== null) searchBT(root, "", answer);
    return answer;
};