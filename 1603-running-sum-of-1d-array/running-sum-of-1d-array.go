func runningSum(nums []int) []int {
    for i, num:=range nums {
        if i==0 {
            continue
        }
        nums[i] = nums[i-1]+num
    }
    return nums
}