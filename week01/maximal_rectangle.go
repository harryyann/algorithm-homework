package week01

func maximalRectangle(matrix [][]byte) int {
	// 想办法凑成84题的样子，一行行遍历，每次都可以构造出一个柱状图数组
	if len(matrix) == 0 || len(matrix[0]) == 0{
		return 0
	}

	// heights记录的是当前行为最底行时的矩形图数组
	var heights = make([]int, len(matrix[0]))
	// 记录答案
	var ans int

	type reg struct{
		width int
		height int
	}

	for _, row := range matrix{
		for j, elem := range row{
			// 当碰到1时，就和上几行中这列的高度累加，否则是0，前面的就都不要了
			if elem == '1'{
				heights[j] += 1
			}else{
				heights[j] = 0
			}
		}
		// 每检查完一行，就可以算一下当前行为最底行时的最大矩形面积，和84题一样了,画个图比较好理解
		heights = append(heights, 0)  // 给heights加个0保证stack可以被排空
		stack := make([]reg, 0)
		var maxArea int
		for _, h := range heights{
			var accumulateWidth int
			// 当不满足单调递增时弹栈，并且计算此时的最大面积，最大面积是这一列的长度乘以累加面宽度
			for len(stack) != 0 && h <= stack[len(stack)-1].height{
				top := stack[len(stack)-1]
				accumulateWidth += top.width
				currentArea := top.height * accumulateWidth
				if maxArea < currentArea{
					maxArea = currentArea
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, reg{height: h, width: accumulateWidth+1})
		}
		// 把heights还原一下，免得数组越加越大
		heights = heights[:len(heights)-1]
		if ans < maxArea{
			ans = maxArea
		}
	}
	return ans
}
