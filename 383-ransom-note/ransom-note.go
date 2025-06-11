func canConstruct(ransomNote string, magazine string) bool {
    m:=make(map[rune]int,0)
    for _,v:=range magazine {
        if _,exist:=m[v]; exist {
            m[v] = m[v]+1
        } else {
            m[v] = 1
        }
    }
    for _,v := range ransomNote {
        if val,exist:=m[v]; exist {
            if val <= 0 {
                return false
            }
            m[v] = m[v]-1
        } else {
            return false
        }
    }
    return true
}