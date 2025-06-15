func reverseBits(num uint32) uint32 {
    var res uint32 = 0
    for i:=0;i<32;i++ {
        x:=num&1
        res=res<<1
        res+=x
        num=num>>1
    }
    return res
}