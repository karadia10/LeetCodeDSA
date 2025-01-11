func removeDuplicates(nums []int) int {
    w:=1
    for r:=0;r<len(nums);r++ {
        if nums[w-1]<nums[r] {
            nums[w] = nums[r]
            w++
        }
    }
    return w
}