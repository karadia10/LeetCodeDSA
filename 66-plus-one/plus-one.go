func plusOne(digits []int) []int {
    carry:=0
    for i:=len(digits) -1 ; i>=0;i-- {
        sum:=digits[i]+1
        if sum==10 {
            carry = 1
            digits[i] = 0
        } else {
            digits[i] = sum
            carry = 0
        }
        if carry == 0 {
            return digits
        }
    }
    if carry == 1 {
        digits = append([]int{1}, digits...)
    }
    return digits
}