func maxDifference(s string) int {
    m:=make(map[rune]int, 0)
    for _,v:=range s {
        if _,exists:=m[v]; !exists {
            m[v] = 1
        } else {
            m[v] = m[v]+1
        }
    }
    max:=0
    min:=100
    for _,v:=range m {
        if v%2 == 0 && v<min {
            min = v
        }
        if v%2 != 0 && v>max {
            max = v
        }
    }
    fmt.Println("min", min)
    fmt.Println("max", max)
    return max-min
}