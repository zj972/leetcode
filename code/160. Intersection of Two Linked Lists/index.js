/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @return {ListNode}
 */
var getIntersectionNode = function(headA, headB) { 
        if(headA === null || headB === null) return null;  
        var p = headA;  
        var q = headB;  
        if(p === q) return p;          
        while(p !== null && q !== null) {  
            p = p.next;  
            q = q.next;  
        }  
        if(p === null) p = headB; else q = headA;  
          
        while(p !== null && q !== null) {  
            p = p.next;  
            q = q.next;  
        }  
        if(p === null) p = headB; else q = headA;  
        while(p !== null && q !== null) {  
            if(p === q) return p;  
            p = p.next;  
            q = q.next;  
        }  
        return null; 
};