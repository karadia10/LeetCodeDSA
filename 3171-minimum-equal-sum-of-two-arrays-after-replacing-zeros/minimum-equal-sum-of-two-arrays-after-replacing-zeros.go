func minSum(nums1 []int, nums2 []int) int64 {
    sum1:=0
    sum2:=0
    count1:=0
    count2:=0
    for _,n := range nums1 {
        sum1+=n
        if (n == 0 ){
            count1+=1
        }
    }
    for _,n := range nums2 {
        sum2+=n
        if (n==0) {
            count2+=1
        }
    }
    // fmt.Println(sum1)
    // fmt.Println(sum2)
    minSum1:= sum1+count1
    minSum2:=sum2+count2
    // fmt.Println("minSum1: ", minSum1)
    // fmt.Println("minSum2: ", minSum2)
    if (minSum1 < minSum2 && count1==0) || (minSum1 > minSum2 && count2 == 0) {
        return -1
    }
    // diff:=minSum1-minSum2
    // fmt.Println("diff: ", diff)
    // if diff < 0 {
    //     return int64(minSum1+diff*-1)
    // }
    // return int64(minSum2+diff)
    return int64(math.Max(float64(minSum1), float64(minSum2)))
    
}