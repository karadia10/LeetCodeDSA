func longestCommonPrefix(strs []string) string {
    ptr :=0
    max:=200
    res:=""
    for _,s:= range strs {
        max = int(math.Min(float64(max), float64(len(s))))
    }
    for ptr < max {
        match := false
        c:=strs[0][ptr]
        for _, s:=range strs {
            if s[ptr] == c {
                match = true
            } else {
                match = false
                break
            }
        }
        if match {
            res+=string(c)
        } else {
            break
        }
        ptr++
    }
    return res
}