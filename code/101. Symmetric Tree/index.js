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
//Distribution 47.13%,runtime 132ms
var isSymmetric = function(root) {
    if(!root) return true;
    var num = 1;
    var rootArray = [root.val];
    function level(root){
        if(root){
            if(rootArray[num] === undefined) rootArray[num] = [];
            if(root.left){
                rootArray[num].push(root.left.val);
            }else{
                rootArray[num].push(null);
            }
            if(root.right){
                rootArray[num].push(root.right.val);
            }else{
                rootArray[num].push(null);
            }
            num++;
            level(root.left);
            level(root.right);
            num--;
        }
    }
    level(root);
    rootArray.length--;
    for(var i = 0, n = rootArray.length; i < n; i++){
        var x = 0;
        var y = rootArray[i].length-1;
        while(x <= y){
            if(rootArray[i][x] !== rootArray[i][y]){
                return false;
            }else{
                x++;
                y--;
            }
        }
    }
    return true;
};
//Distribution 16.09%,runtime 144ms
var isSymmetric = function(root) {
    function isSymmetricHelp(left,right){
        if(left === null || right === null)
            return left === right;
        if(left.val !== right.val)
            return false;
        return isSymmetricHelp(left.left, right.right) && isSymmetricHelp(left.right, right.left);
    }
    return root === null || isSymmetricHelp(root.left, root.right);
};