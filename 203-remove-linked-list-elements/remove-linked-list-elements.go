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
            // fmt.Println("curr: ",curr)
            curr.Next = curr.Next.Next
            // fmt.Println("curr next", curr.Next)
            // fmt.Println("curr updated: ",curr)
        } else{
            curr = curr.Next
        }
        // fmt.Println("curr new: ", curr)
    }
    return tmp.Next
}