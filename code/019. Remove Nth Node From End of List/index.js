/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
var removeNthFromEnd = function(head, n) {
    var last = head;
    var p = head;
    for(var i = 0; i < n; i++){
        if(last.next === null){
            head = head.next;
            return head;
        }else{
            last = last.next;
        }
    }
    while(last.next !== null){
        p = p.next;
        last  = last.next;
    }
    var pNext = p.next;
    p.next = pNext.next;
    return head;
};