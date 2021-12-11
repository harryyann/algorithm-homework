package week08

// https://leetcode-cn.com/problems/number-of-islands

var gridArr [][]byte
var visited [][]bool
var m, n int

func numIslands(grid [][]byte) int {
	if len(grid) == 0{
		return 0
	}
	gridArr = grid
	m = len(grid)
	n = len(grid[0])
	visited = make([][]bool, m)
	for v, _ := range visited{
		visited[v] = make([]bool, n)
	}
	var ans int
	for i:=0; i<m; i++{
		for j:=0; j<n; j++{
			if grid[i][j] == '1' && !visited[i][j]{
				ans += 1
				dfs(i, j)
			}
		}
	}
	return ans
}

func dfs(x, y int){
	if x >= m || x < 0 || y >= n || y < 0 || gridArr[x][y] != '1' || visited[x][y]{
		return
	}
	visited[x][y] = true
	dfs(x, y + 1)
	dfs(x + 1, y)
	dfs(x, y - 1)
	dfs(x - 1, y)
}
