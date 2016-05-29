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
var isPalindrome = function(head) {
    var list = [];
    while(head){
        list.push(head.val);
        head = head.next;
    }
    for(var i = 0; i < list.length; i++){
        if(list[i] !== list[list.length-1-i]) return false;
    }
    return true;
};