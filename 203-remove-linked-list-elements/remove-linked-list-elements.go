/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    tmp:=&ListNode{Next:head}
    curr:=tmp
    for curr != nil && curr.Next != nil {
        if curr.Next.Val == val {
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }
    return tmp.Next
}