/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {boolean}
 */
//Distribution 16.67%,runtime 124ms
var hasCycle = function(head) {
    if(!head || !head.next) return false;
    var n = head, m = head;
    while(1){
        if(n && n.next){
            n = n.next;
        }else{
            return false;
        }
        if(m && m.next && m.next.next){
            m = m.next.next;
        }else{
            return false;
        }
        if(n === m) return true;
    }
    return false;
};