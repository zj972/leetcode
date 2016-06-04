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
//Distribution 18.75%,runtime 162ms
var oddEvenList = function(head) {
    if(!head) return head;
    var ohead = head.next;
    var j = head;
    var o = head.next;
    var node = head.next;
    while(node && node.next){
        j.next = node.next;
        o.next = node.next.next;
        j = j.next;
        o = o.next;
        node = j.next;
    }
    j.next = ohead;
    return head;
};