/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
 //迭代
//Distribution 84.62%,runtime 104ms
var reverseList = function(head) {
  var pre = null;
  while (head) {
    var next = head.next;
    head.next = pre;
    pre = head;
    head = next;
  }
  return pre;
};
//递归
//Distribution 18.75%,runtime 124ms
var reverseList = function(head) {
  if (head === null || head.next === null)
    return head;

  var next = head.next;
  head.next = null;
  var newHead = reverseList(next);
  next.next = head;

  return newHead;
};