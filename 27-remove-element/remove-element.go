func removeElement(nums []int, val int) int {
    m:=make(map[int]int, 0)
    for _,v := range nums {
        if (v == val) {
            continue
        }
        if _,exist:=m[v]; exist {
            m[v]=m[v]+1
        } else {
            m[v]=0
        }
    }
    p:=0
    for k,v:=range m {
        for i:=0;i <= v;i++ {
            nums[p]=k
            p++
        }
    }
    return p
}
