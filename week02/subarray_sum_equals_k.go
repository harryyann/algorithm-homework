package week02

// https://leetcode-cn.com/problems/subarray-sum-equals-k/

// 方法1
func subarraySum1(nums []int, k int) int {
	var ans int
	for i:=0; i<len(nums); i++{
		var accumulate int
		for j:=i; j>=0; j--{
			accumulate = accumulate + nums[j]
			if accumulate == k{
				ans++
			}
		}
	}
	return ans
}
