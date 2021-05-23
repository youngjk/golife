/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var h1, h2 *ListNode = l1, l2
    resultHeadNode := ListNode{Val: 0, Next: nil}
    current := &resultHeadNode
    carryover := 0
    
    for (h1 != nil || h2 != nil) {
        v1, v2 := 0, 0
        if (h1 != nil) { v1 = h1.Val }
        if (h2 != nil) { v2 = h2.Val }
        
        sum := carryover + v1 + v2
        carryover = sum/10
        
        current.Next = &ListNode{Val: sum%10, Next: nil}
        current = current.Next
        
        if (h1 != nil) { h1 = h1.Next }
        if (h2 != nil) { h2 = h2.Next }
    }
    
    if (carryover>0) {
        current.Next = &ListNode{Val: carryover, Next: nil}
    }
    
    return resultHeadNode.Next
}