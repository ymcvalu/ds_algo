package algorithm

// 求数组逆序对数量
// 可以在归并排序过程中计算逆序对数量
// arr[i,j] => left:arr[i,mid) || right:arr[mid,j]
// 现在假设left和right都已经是有序的数组了
// 在归并过程中，如果存在right中的元素小于left中的元素，则记录逆序对数量

func InversePairs(nums []int) int {
    ln:=len(nums)
    if ln<2{
        return 0
    }

    mid := ln/2
    left := nums[:mid]
    right := nums[mid:]
    n := InversePairs(left)
    n += InversePairs(right)
    l,r := 0,0
    ll, lr := len(left), len(right)

    buf := make([]int, 0, ln)
    for l < ll && r < lr{
        if left[l]<=right[r]{
            buf=append(buf,left[l])
            l++
        }else {
            buf=append(buf,right[r])
            n+=ll-l
            r++
        }
    }

    if l<ll{
        buf=append(buf, left[l:]...)
    }else {
        buf=append(buf,right[r:]...)
    }

    copy(nums,buf)
    return n
}

// 可以优化，预分配归并排序过程中需要用到的临时切片
func reversePairs(nums []int)int{
    if len(nums)<2{
        return 0
    }
    // 预分配临时切片
    buf:=make([]int,len(nums))
    return  _reversePairs(nums,buf)
}

func _reversePairs(nums []int,buf []int) int {
    ln:=len(nums)
    if ln<2{
        return 0
    }

    mid := ln/2
    left := nums[:mid]
    right := nums[mid:]
    n := _reversePairs(left,buf[:mid])
    n += _reversePairs(right,buf[mid:])
    l,r := 0,0
    ll, lr := len(left), len(right)
    buf=buf[:0]
    for l < ll && r < lr{
        if left[l]<=right[r]{
            buf=append(buf,left[l])
            l++
        }else {
            buf=append(buf,right[r])
            n+=ll-l
            r++
        }
    }

    if l<ll{
        buf=append(buf, left[l:]...)
    }else {
        buf=append(buf,right[r:]...)
    }

    copy(nums,buf)
    return n
}