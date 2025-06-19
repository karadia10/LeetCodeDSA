func reverseWords(s string) string {
    s = strings.TrimSpace(s)
    arr:=strings.Fields(s)
    for i,j:=0, len(arr)-1; i<j; i,j=i+1,j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
    return strings.Join(arr, " ")
}