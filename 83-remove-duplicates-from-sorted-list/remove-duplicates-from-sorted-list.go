/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    curr:=head
    for curr != nil && curr.Next!=nil {
        if  curr.Next.Val == curr.Val {
            curr.Next = curr.Next.Next
        } else {
            curr=curr.Next
        }
    }
    return head
}