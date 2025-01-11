func lengthOfLastWord(s string) int {
    res:=0
    l:=len(s) - 1
    for l>=0 {
        if s[l] != ' ' {
            res++
        } else if res > 0 {
            return res
        }
        l--
    }
    return res
}