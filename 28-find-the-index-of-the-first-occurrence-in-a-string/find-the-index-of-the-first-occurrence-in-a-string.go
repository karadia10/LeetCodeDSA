func strStr(haystack string, needle string) int {
    if len(haystack) < len(needle) {
        return -1
    }
    for i:=0;i<len(haystack);i++ {
        if (haystack[i]==needle[0]) && len(haystack)-i>=len(needle) {
            found:=true
            for j:=0;j<len(needle);j++ {
                fmt.Println("i:",i," j:",j)
                if haystack[i+j]!=needle[j] {
                    found=false
                    break
                } else {
                    found = true
                }
            }
            if found {
                return i
            }
        }
    }
    return -1
}