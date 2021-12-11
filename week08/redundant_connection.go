package week08

// https://leetcode-cn.com/problems/redundant-connection/

var visited []bool
var to  [][]int
var hasCycle bool
var n int
func findRedundantConnection(edges [][]int) []int {
	for _, edge := range edges{
		x := edge[0]
		y := edge[1]
		if x > y{
			if x > n{
				n = x
			}
		}else{
			if y > n{
				n = y
			}
		}
	}

	to = make([][]int, n+1)
	visited = make([]bool, n+1)

	for _, edge := range edges{
		x := edge[0]
		y := edge[1]
		to[x] = append(to[x], y)
		to[y] = append(to[y], x)

		// 复原现场
		hasCycle = false
		visited = make([]bool, n+1)
		// 每一次加边，都进行一次深度递归，加到哪一个边hasCycle变为true了就找到了这个冗余的连接
		dfs(x, 0)
		if hasCycle{
			return edge
		}
	}
	return nil
}

// 接收每个节点，并逐层遍历
func dfs(x int, par int){
	visited[x] = true
	for _, y := range to[x]{
		if par == y{
			// 如果当前的节点和上一个节点相同的话，就要跳过
			continue
		}
		// 如果x能到的节点y没走过，则继续dfs y
		if !visited[y]{
			dfs(y, x)
		}else{
			// 如果y走过了，并且y不是x的上一个节点，则这个连接就是要找的冗余连接
			hasCycle = true
			return
		}
	}
	return
}
