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
var swapPairs = function(head) {
    if(!head) return head;
    var newList = head;
    while(1){
        if(head && head.next){
            var temp = head.val;
            head.val = head.next.val;
            head.next.val = temp;
            if(head.next.next){
                head = head.next.next;
            }else{
                return newList;
            }
        }else if(head){
            return newList;
        }
    }
};