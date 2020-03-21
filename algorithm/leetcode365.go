package algorithm

/**
有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？

如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。

你允许：

装满任意一个水壶
清空任意一个水壶
从一个水壶向另外一个水壶倒水，直到装满或者倒空
*/

// ax+by=z
// 贝祖定理：z是x和y的最大公约数的倍数
func canMeasureWaterMath(x, y, z int) bool {
	if z == 0 || z == x+y || z == x || z == y {
		return true
	}
	if z > x+y {
		return false
	}

	return z%gcd(x, y) == 0
}

type ContainerState struct {
	x int
	y int
}

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 || z == x+y || z == x || z == y {
		return true
	}
	if z > x+y {
		return false
	}

	state := ContainerState{
		x: 0,
		y: 0,
	}

	return canMeasureWaterBT(state, x, y, z, map[ContainerState]bool{})
}

func canMeasureWaterBT(state ContainerState, x, y int, z int, filter map[ContainerState]bool) bool {
	if state.x == z || state.y == z || state.x+state.y == z {
		return true
	}

	next := func(ns ContainerState) bool {
		if !filter[ns] {
			filter[ns] = true
			return canMeasureWaterBT(ns, x, y, z, filter)
		}
		return false
	}

	if state.x != x {
		ns := state
		ns.x = x
		if next(ns) {
			return true
		}
	}

	if state.y != y {
		ns := state
		ns.y = y
		if next(ns) {
			return true
		}
	}

	if state.x > 0 && state.y != y {
		ns := state
		ns.y += ns.x
		if ns.y > y {
			ns.x = ns.y - y
			ns.y = y
		} else {
			ns.x = 0
		}
		if next(ns) {
			return true
		}
	}

	if state.y > 0 && state.x != x {
		ns := state
		ns.x += ns.y
		if ns.x > x {
			ns.y = ns.x - x
			ns.x = x
		} else {
			ns.y = 0
		}
		if next(ns) {
			return true
		}
	}

	if state.x > 0 {
		ns := state
		ns.x = 0
		if next(ns) {
			return true
		}
	}

	if state.y > 0 {
		ns := state
		ns.y = 0
		if next(ns) {
			return true
		}
	}

	return false
}
