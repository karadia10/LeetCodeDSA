func climbStairs(n int) int {
    one:=1
    two:=0
    for n>0 {
        one, two = one+two, one
        n--
    }
    return one

}