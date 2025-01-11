func removeElement(nums []int, val int) int {
    w:=0
    for r:=0;r<len(nums);r++ {
        if nums[r]!=val {
            nums[w] = nums[r]
            w++
        }
    }
    return w
}
