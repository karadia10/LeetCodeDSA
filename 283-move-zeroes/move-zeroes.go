//working code
func moveZeroes(nums []int)  {
    ptr1:=0
    ptr2:=1
    for ptr2<len(nums) && ptr1<len(nums) {
        if nums[ptr1] !=0 {
            ptr1++
            ptr2=ptr1+1
            continue
        }
        if nums[ptr2] != 0 && nums[ptr1] == 0 {
            swap(&nums, ptr1, ptr2)
            ptr1++
        }
        ptr2++
    }
}

func swap(nums *[]int,a, b int) {
    tmp:=(*nums)[a]
    (*nums)[a] = (*nums)[b]
    (*nums)[b] = tmp
}