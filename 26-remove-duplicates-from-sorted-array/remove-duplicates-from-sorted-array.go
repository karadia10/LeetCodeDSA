func removeDuplicates(nums []int) int {
    w:=0
    for r:=0;r<len(nums);r++ {
        if nums[w]<nums[r] {
            nums[w+1] = nums[r]
            w++
        }
    }
    return w+1
}