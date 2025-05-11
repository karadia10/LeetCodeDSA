/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    root := &ListNode{Val: 0}
    ptr := root
    for l1 != nil || l2 !=nil || carry != 0 {
        s:= carry
        if l1!=nil {
            s+=l1.Val
            l1 = l1.Next
        }
        if l2!=nil {
            s+=l2.Val
            l2 = l2.Next
        }
        carry = s/10
        ptr.Next = &ListNode{Val: s % 10}
        ptr = ptr.Next 
    }
    return root.Next
}