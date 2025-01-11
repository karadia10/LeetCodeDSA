func longestCommonPrefix(strs []string) string {
    //Approach 1: O(n^2)
    // ptr :=0
    // max:=200
    // res:=""
    // for _,s:= range strs {
    //     max = int(math.Min(float64(max), float64(len(s))))
    // }
    // for ptr < max {
    //     match := false
    //     c:=strs[0][ptr]
    //     for _, s:=range strs {
    //         if s[ptr] == c {
    //             match = true
    //         } else {
    //             match = false
    //             return res
    //         }
    //     }
    //     if match {
    //         res+=string(c)
    //     }
    //     ptr++
    // }
    // return res
    //Approach 2:O(nlog(n))
    if len(strs) == 1 {
        return strs[0]
    }
    //if there is a difference the first and last will have the most difference
    sort.Strings(strs)
    l:=len(strs)
    //iterate on first string in sorted array
    for i := range strs[0] {
        if strs[0][i] != strs[l-1][i] {
            return strs[0][:i]
        }
    }
    return strs[0]
}