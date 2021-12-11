package week04

// https://leetcode-cn.com/problems/surrounded-regions/

func solve(board [][]byte)  {
	// 不被X包围的O一定是在边界，或者和边界上的某个O相连
	// 则该题就是找到边界上的O，然后用dfs找到所有边界上的O块，将除了这些O块以外的O变为X即可
	if len(board) == 0 || len(board[0]) == 0{
		return
	}
	m := len(board)  //行数
	n := len(board[0]) // 列数

	// 遍历第一行
	for i:=0; i<n; i++{
		dfs(0, i, m, n, board)
	}
	// 遍历最后一行
	for i:=0; i<n; i++{
		dfs(m-1, i, m, n, board)
	}

	// 遍历第一列
	for i:=0; i<m; i++{
		dfs(i, 0, m, n, board)
	}

	// 遍历最后一列
	for i:=0; i<m; i++{
		dfs(i, n-1, m, n, board)
	}
	// 最后再遍历，把所有O改为X，把所有-改为O即可
	for i:=0; i<m; i++{
		for j:=0; j<n; j++{
			if board[i][j] == 'O'{
				board[i][j] = 'X'
			}
			if board[i][j] == '-'{
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(i, j int, m, n int, board [][]byte){
	if i == m || i < 0 ||  j == n || j < 0{
		return
	}
	if board[i][j] != 'O'{
		return
	}
	// 每个i, j以及与之相连的为O的都改为'-'
	(board)[i][j] = '-'
	dfs(i+1, j, m, n, board)
	dfs(i-1, j, m, n, board)
	dfs(i, j+1, m, n, board)
	dfs(i, j-1, m, n, board)
}