func countTriplets(arr []int64, r int64) int64 {
    m:=map[int64]int{}
    for _,k:=range arr{
        if _,ok:=m[k];ok{
            m[k]++
        }else{
            m[k]=1
        }
    }
    var result []int64 
    for k,v:=range m{
        val,ok:=m[r*k]
        val2,sk:=m[r*r*k]
        if ok && sk{
            append(int64(v*val*val2), result)
        }
    }
    var ans int64
    for _,v:=range result{
        ans+=v
    }
    return ans
}
