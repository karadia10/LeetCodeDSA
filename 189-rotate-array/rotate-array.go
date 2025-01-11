func rotate(nums []int, k int)  {
    len:=len(nums)
    k=k%len
    reverse(&nums, 0, len-1)
    reverse(&nums, 0, k-1)
    reverse(&nums, k, len-1)
}

func reverse(nums *[]int, l,r int) {
    for l<r {
        (*nums)[l], (*nums)[r] = (*nums)[r], (*nums)[l]
        l++
        r--   
    }
    fmt.Println(nums)
}