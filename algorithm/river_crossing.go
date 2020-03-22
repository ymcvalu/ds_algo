package algorithm

/**
问题描述：
有三个传教士和三个吃人恶魔要渡过一条河，河中有一条船，只能装下两个人。在任何地方（无论是岸边还是船上），如果吃人恶魔数量多于传教士数量，吃人恶魔就会吃掉传教士。问：怎么才能让这些都安全过河？

解题思路：
首先定义状态：
(L1, L2, B, R1, R2)
• L1：左岸传教士人数
• L2：左岸吃人恶魔人数
• B：当前船是靠在左岸还是右岸，取值可以为L或者R
• R1：右岸传教士人数
• R2：右岸吃人恶魔人数

首先，传教士和恶魔都可以通过船来回两岸，实际上这是一道与图论相关的题目。
状态可以分为两类：
• B为L：船靠在左岸，传教士和恶魔可以坐船到右岸
• B为R：船靠在右岸，传教士和恶魔可以坐船到左岸
B为L的状态可以转换为B为R，反之也可，因此所有状态实际是二分图。
*/
type CrossRiverState struct {
	L1     int
	L2     int
	AtLeft bool
	R1     int
	R2     int
}

func river_crossing(state CrossRiverState, path []CrossRiverState, filter map[CrossRiverState]bool, result *[][]CrossRiverState) {
	if !state.AtLeft && state.L1 == 0 && state.L2 == 0 {
		dup := make([]CrossRiverState, len(path))
		copy(dup, path)
		*result = append(*result, dup)

		return
	}

	atLeft := state.AtLeft
	if !atLeft {
		state.L1, state.R1 = state.R1, state.L1
		state.L2, state.R2 = state.R2, state.L2
	}

	traver := func(s CrossRiverState) {
		if !atLeft {
			s.L1, s.R1 = s.R1, s.L1
			s.L2, s.R2 = s.R2, s.L2
		}
		s.AtLeft = !atLeft

		_s := s
		_s.AtLeft = true // 禁止同一条边重复往返
		// 环检测
		if filter[_s] {
			return
		}

		path = append(path, s)

		filter[_s] = true
		river_crossing(s, path, filter, result)
		path = path[:len(path)-1]
		delete(filter, _s)
	}

	for i := 1; i <= state.L1 && i <= 2; i++ {
		for j := 1; j <= state.L2 && j+i <= 2; j++ {
			// i个传教士，j个吃人恶魔过河
			if i >= j && state.R1+i >= state.R2+j && state.L1-i >= state.L2-j {

				newState := CrossRiverState{
					L1: state.L1 - i,
					L2: state.L2 - j,
					R1: state.R1 + i,
					R2: state.R2 + j,
				}

				traver(newState)
			}
		}
	}

	// i个传教士过河
	for i := 1; i <= state.L1 && i <= 2; i++ {
		if (state.L1-i >= state.L2 || state.L1 == i) && state.R1+i >= state.R2 {
			newState := CrossRiverState{
				L1: state.L1 - i,
				L2: state.L2,
				R1: state.R1 + i,
				R2: state.R2,
			}
			traver(newState)

		}
	}

	// j个吃人恶魔过河
	for j := 1; j <= state.L2 && j <= 2; j++ {
		if state.R1 == 0 || state.R2+j <= state.R1 {
			newState := CrossRiverState{
				L1: state.L1,
				L2: state.L2 - j,
				R1: state.R1,
				R2: state.R2 + j,
			}
			traver(newState)
		}
	}

	// 允许空船返回
	if !atLeft {
		empty := CrossRiverState{
			AtLeft: true,
			L1:     state.R1,
			L2:     state.R2,
			R1:     state.L1,
			R2:     state.L2,
		}

		path = append(path, empty)
		river_crossing(empty, path, filter, result)
		path = path[:len(path)-1]
	}

}
