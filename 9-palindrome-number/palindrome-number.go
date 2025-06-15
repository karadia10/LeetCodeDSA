func isPalindrome(x int) bool {
    rev:=0
    num:=x
    for num>0 {
        x:=num%10
        num/=10
        rev=rev*10+x
    }
    return rev == x
}