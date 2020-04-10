package algorithm

/**
将int32的范围[l,r]的bit设置为1
0表示第1位
*/
func setBitByRange(n int32, l, r int) int32 {
	if l < r {
		panic("invalid range")
	}
	if r < 0 {
		r = 0
	}
	if l > 31 {
		l = 31
	}

	var mask1 uint32 = (1 << uint(r)) - 1
	var mask2 uint32 = ^((1 << uint(l+1)) - 1)
	mask := ^((mask1 | mask2) & 0xffffffff)
	return n | int32(mask)
}
