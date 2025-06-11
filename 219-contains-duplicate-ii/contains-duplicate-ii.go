func containsNearbyDuplicate(nums []int, k int) bool {
    m:=make(map[int]int, 0)
    for i,v:=range nums {
        if val,exists:=m[v]; exists {
            if i-val <= k {
                return true
            }
        }
        m[v]=i
    }
    return false
}