package week02

// https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/submissions/

// 超出时间限制了
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0{
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])

	// 前缀和数组
	var preSum = make([][]int, m+1)
	for i, _ := range preSum{
		preSum[i] = make([]int, n+1)
	}

	// k:前缀和的值，v:计数
	ansMap := make(map[int]int)

	//先套公式求前缀和
	for i:=1; i<=m; i++{
		for j:=1; j<=n; j++{
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] -preSum[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	// 暴力枚举可能的
	for i:=1; i<=m; i++{
		for j:=1; j<=n; j++{
			for p:=i; p>0; p--{
				for q:=j; q>0; q--{
					// 计算从(p, q)到(i, j)之间的子矩阵和，存入ansMap
					sum := preSum[i][j] - preSum[i][q-1] - preSum[p-1][j] + preSum[p-1][q-1]
					ansMap[sum] ++
				}
			}
		}
	}
	return ansMap[target]
}
