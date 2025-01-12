func twoSum(numbers []int, target int) []int {
    m:=make(map[int]int, 0)
    for i,v:=range numbers {
        _,exist:=m[target-v]
        if !exist {
            m[target-v]=i
        }
    }
    arr:=make([]int,0)
    for i,v :=range numbers {
        val,exist:=m[v]
        if exist && val!=i {
            if val>i {
                arr=append(arr, i+1, val+1)
            } else {
                arr=append(arr, val+1, i+1)
            }
            break
        }
    }
    return arr
}