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
var deleteDuplicates = function(head) {
    if(!head || !head.next) return head;
    var n = head;
    while(n.next){
        if(n.val === n.next.val){
            var nextNode = n.next;
            n.next = nextNode.next;
        }else{
            n = n.next;
        }
    }
    return head;
};