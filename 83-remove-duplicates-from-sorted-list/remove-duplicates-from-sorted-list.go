/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    curr:=head
    for curr != nil {
        if curr.Next!=nil && curr.Next.Val == curr.Val {
            // fmt.Println("1")
            if curr.Next.Next != nil {
                // fmt.Println("2", curr.Next.Next)
                curr.Next = curr.Next.Next
            } else {
                // fmt.Println("3") 
                curr.Next=nil
            }
            continue
        }
        fmt.Println(curr)
        curr=curr.Next
    }
    return head
}