func isPalindrome(s string) bool {
    l:=0
    r:=len(s)-1
    for l<=r {
        // fmt.Println(l, r)
        if !isAlphabet(s[l]){
            l++
            continue
        }
        if !isAlphabet(s[r]) {
            r--
            continue
        }
        if (unicode.ToUpper(rune(s[l])) != unicode.ToUpper(rune(s[r]))) {
            return false
        } 
        l++
        r--
    }
    return true
}
func isAlphabet(a byte) bool {
    return ((a>='a' && a<='z') || (a>='A' && a<='Z')) || (a>='0' && a<='9')
}