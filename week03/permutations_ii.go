package week03

import "sort"

// https://leetcode-cn.com/problems/permutations-ii/

func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	scheme:= []int{}
	used := map[int]bool{}
	// 一定要先排序，dfs里的剪枝条件才成立
	sort.Ints(nums)
	dfs(0, nums, scheme, used, &ans)
	return ans
}

func dfs(i int, nums []int, scheme []int, used map[int]bool, ans *[][]int){
	if i == len(nums){
		temp := make([]int, len(scheme))
		copy(temp, scheme)
		*ans = append(*ans, temp)
	}

	for j:=0; j<len(nums); j++{
		// 要判断排列是不是有重复的话，
		if !used[j]{
			// j > 0 防止数组越界，因为等于0是肯定不重复，判断这个元素是否和前一个相同，相同就要被剪掉
			// 这里用used[j-1]==true或used[j-1]==false都可以，false是剪大枝，true是剪小枝
			if j > 0 && nums[j] == nums[j-1] && !used[j-1]{
				continue
			}
			used[j] = true
			scheme = append(scheme, nums[j])
			dfs(i+1, nums, scheme, used, ans)
			scheme = scheme[:len(scheme)-1]
			used[j] = false
		}
	}
}
