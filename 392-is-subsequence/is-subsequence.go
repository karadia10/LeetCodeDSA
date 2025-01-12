func isSubsequence(s string, t string) bool {
    p1:=0
    if len(s) == 0 {
        return true
    }
    for i:= range t {
        if t[i] == s[p1] {
            p1++
        }
        if p1 >= len(s) {
            return true
        }
    }
    return false
}