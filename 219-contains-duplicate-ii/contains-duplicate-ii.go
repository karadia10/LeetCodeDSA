func containsNearbyDuplicate(nums []int, k int) bool {
    m:=make(map[int]int, 0)
    for i,v:=range nums {
        if _,exists:=m[v]; exists {
            if math.Abs(float64(i-m[v])) <= float64(k) {
                return true
            }
        }
        m[v]=i
    }
    return false
}