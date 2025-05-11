func threeConsecutiveOdds(arr []int) bool {
    count:=0
    for _,n := range arr {
        if n%2 == 0 {
            count = 0
        } else {
            count++
        }
        if count == 3 {
            return true
        } 
    }
    return false
}