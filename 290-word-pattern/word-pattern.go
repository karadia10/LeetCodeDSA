func wordPattern(pattern string, s string) bool {
    m:=make(map[rune]string, 0)
    p:=make(map[string]rune,0)
    str:=strings.Split(s, " ")
    if len(pattern) != len(str) {
        return false
    }
    for i,v:=range pattern {
        if _,exist:=m[v]; exist {
            if str[i] != m[v]  {
                return false
            }
        } else {
            m[v] = str[i]
            if val,ex:=p[str[i]]; ex && val != v {
                return false
            } else {
                p[str[i]]=v
            }
           
        }
    }
    fmt.Println(m)
    return true
}