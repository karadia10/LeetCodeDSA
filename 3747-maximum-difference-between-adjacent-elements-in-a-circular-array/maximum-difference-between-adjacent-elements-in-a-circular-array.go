func maxAdjacentDistance(nums []int) int {
    max:=math.Abs(float64(nums[0]-nums[len(nums)-1]))

    for i:=range nums {
        if i==len(nums)-1 {
            break
        }
        curr:=math.Abs(float64(nums[i]-nums[i+1]))
        if  curr > max {
            max = curr 
        }
    }
    return int(max)
}