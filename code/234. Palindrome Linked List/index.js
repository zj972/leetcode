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
// 反转后一半链表对比
// var isPalindrome = function(head) {
//     if (head === null) {  
//         return true;  
//     }  
//     var slow = head;  
//     var fast = slow.next;  
//     while (fast !== null && fast.next !== null && fast.next.next !== null){  
//         slow = slow.next;  
//         fast = fast.next.next;  
//     }  
//     var p = slow.next;  
//     var q;  
//     var end = null;  
//     while (p !== null) {  
//         q = p.next;  
//         p.next = end;  
//         end = p;  
//         p = q;  
//     }  
//     while (head !== null && end !== null) {  
//         if (head.val != end.val) {  
//             return false;  
//         }  
//         head = head.next;  
//         end = end.next;  
//     }  
//     return true;  
// };