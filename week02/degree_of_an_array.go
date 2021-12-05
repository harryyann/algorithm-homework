package week02

// https://leetcode-cn.com/problems/degree-of-an-array/

func findShortestSubArray(nums []int) int {
	// 键是元素值，值是elemInfo对象，记录这个元素出现的最小索引和最大索引
	type elemInfo struct{
		du int
		smallIndex int
		bigIndex int
	}
	var ans = make(map[int]*elemInfo)
	var du int
	for i, n := range nums{
		info, has := ans[n]
		if has{
			info.du += 1
			if info.smallIndex > i{
				info.smallIndex = i
			}
			if info.bigIndex < i {
				info.bigIndex = i
			}
		}else{
			var info = elemInfo{}
			info.smallIndex = i
			info.bigIndex = i
			ans[n] = &info
			info.du += 1
		}
		if du < ans[n].du{
			du = ans[n].du
		}
	}
	var finalAns int
	for k, _ := range ans{
		if ans[k].du == du{
			length := ans[k].bigIndex - ans[k].smallIndex + 1
			if finalAns == 0{
				finalAns = length
			}else{
				if finalAns > length{
					finalAns = length
				}
			}
		}
	}
	return finalAns
}