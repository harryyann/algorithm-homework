package week09

import "strings"

// https://leetcode-cn.com/problems/jewels-and-stones

func numJewelsInStones(J string, S string) int {
	count := 0
	for i := range S {
		if strings.Contains(J, string(S[i])) {
			count++
		}
	}
	return count
}
