/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    return same(p,q)
}

func same(pnode *TreeNode, qnode *TreeNode) bool {
    // if pnode == nil && qnode == nil {
    //     return true
    // }
    if pnode == nil && qnode == nil {
        return true
    } else if pnode ==nil || qnode == nil {
        return false
    } 
    if pnode.Val != qnode.Val {
        return false
    }
    return same(pnode.Left, qnode.Left) && same(pnode.Right, qnode.Right)
}