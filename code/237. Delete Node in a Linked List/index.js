/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} node
 * @return {void} Do not return anything, modify node in-place instead.
 */
//Distribution 65.34%,runtime 132ms
var deleteNode = function(node) {
    var nextNode = node.next;
    node.val = nextNode.val;
    node.next = nextNode.next;
};