import "fmt"
func merge(nums1 []int, m int, nums2 []int, n int)  {
    ans:=make([]int,0)
    p1:=0
    p2:=0
    
    for p1+p2<=m+n {
        if p1>=m {
            ans=append(ans, nums2[p2:n]...)
            break
        } else if p2>=n {
            ans=append(ans, nums1[p1:m]...)
            break
        }
        if nums1[p1]>nums2[p2] {
            ans=append(ans,nums2[p2])
            p2++
        } else if nums1[p1]<nums2[p2]{
            ans=append(ans,nums1[p1])
            p1++
        } else {
            ans = append(ans,nums1[p1])
            p1++
        }
    }

    for i,v:=range ans {
        nums1[i]=v
    }
}