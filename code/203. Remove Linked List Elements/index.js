/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} val
 * @return {ListNode}
 */
var removeElements = function(head, val) {
    while(head){
        if(head.val === val)
            head = head.next;
        else break;
    }
    if(!head) return head;
    var list = head;
    while(list.next){
        if(list.next.val === val){
            list.next = list.next.next;
        }else{
            list = list.next;
        }
    }
    return head;
};