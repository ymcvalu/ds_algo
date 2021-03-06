package algorithm

// n层楼，k个鸡蛋，存在楼层f，从[1,f)层往下扔鸡蛋不会碎，[f,n]扔都会碎
// 如果一个鸡蛋扔下去后没有碎，可以捡起来再扔
// 现在给定n和k，求出一种寻找策略，通过这个策略存在楼层f，最坏情况下所需要扔鸡蛋的次数最少

// 具体楼层f是哪一层是未知的，可能是任意一层
// 通过某种策略来查找楼层f时，根据f在[1..n]的不同，具体抛鸡蛋的次数记为集合T=[T1,T2,...,Tn]
// 所谓最坏情况，就是集合T中的最大值Tmax
// 我们要寻找一种策略，使得Tmax总是最小的

// 如果鸡蛋的个数不限，那么最优的策略是通过二分法
// 但是鸡蛋个数是有限的，如果k=1，那么只能从第一楼开始往上遍历
// 如果鸡蛋有2个，那么可以先尝试50楼，如果碎了，第二个鸡蛋从第1楼开始遍历，如果没有碎，从51楼开始遍历
// 如果有三个鸡蛋...
// ...

// 假设函数f(n,k)表示n层楼，k个鸡蛋的最优解
// 当我们尝试从楼层h开始抛鸡蛋：
//   1. 如果鸡蛋碎了，则总共需要抛1+f(h-1,k-1)
//   2. 如果鸡蛋没有碎，则总共需要抛1+f(n-h,k)，抛鸡蛋的次数只与楼层数量有关，与当前处于第几层并没有关系
// 那么从楼层h开始抛鸡蛋，最坏情况就是max(f(h-1,k-1), f(n-h,k)) + 1，记为Th
// 而如果从1..n遍历h，那么f(n,k)=min(T1,T2,...,Tn)

// f(0,k) = 0 (k>=1)
// f(n,1) = n (n>=1)
// f(n,m) = min{ max{ f(h-1,m-1), f(n-h,m) } for h in 1..n } + 1

func throwEgg(n, k int) int {
	if k == 0 || n == 0 {
		return 0
	}
	if k == 1 {
		return n // 只有一个鸡蛋，只能从第一层开始遍历，最坏情况需要扔n次鸡蛋
	}

	m := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		m[i] = make([]int, n+1)
		if i == 0 {
			continue
		}
		// 只有一个鸡蛋的时候
		for j := 0; j < n+1; j++ {
			m[i][j] = j
		}
	}

	for ki := 2; ki <= k; ki++ {
		for ni := 1; ni <= n; ni++ {
			for h := 1; h <= ni; h++ {
				t := max(m[ki][ni-h], m[ki-1][h-1]) + 1
				m[ki][ni] = min(m[ki][ni], t)
			}
		}
	}

	return m[k][n]
}
