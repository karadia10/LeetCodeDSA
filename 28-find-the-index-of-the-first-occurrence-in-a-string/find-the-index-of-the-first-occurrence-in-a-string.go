func strStr(haystack string, needle string) int {
    if len(haystack) < len(needle) {
        return -1
    }
    for i:=0;i<=(len(haystack)-len(needle));i++ {
        if (haystack[i]==needle[0]) {
            found:=true
            for j:=0;j<len(needle);j++ {
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