func majorityElement(nums []int) int {
    m:=make(map[int]int,0)
    for _,v:=range nums {
        _,exist:=m[v]
        if exist {
            m[v]+=1
        } else {
            m[v]=1
        }
        if m[v]>len(nums)/2 { 
            return v
        }
    }
    return 0
}