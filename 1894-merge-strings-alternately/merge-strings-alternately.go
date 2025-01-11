func mergeAlternately(word1 string, word2 string) string {
    str:=""
    l1:=len(word1)
    l2:=len(word2)
    for i := range int(math.Max(float64(l1),float64(l2))) {
        if i<l1 {
            str += string(word1[i])
        }
        if i<l2 {
            str += string(word2[i])
        }
    }
    return str
}