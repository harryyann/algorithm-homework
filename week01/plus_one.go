package week01

// https://leetcode-cn.com/problems/plus-one/

func plusOne(digits []int) []int {
	if len(digits) == 0{
		return []int{}
	}
	for i:=len(digits)-1; i>=0; i-- {
		// 如果数字不是9，给这个数字加1返回即可
		if digits[i] != 9{
			digits[i] ++
			return digits
		}
		// 如果数字是9， 那么要把这一位置为0，并且继续循环前一位
		digits[i] = 0
	}
	// 如果没有在循环体内返回，证明数字全是9，那就需要把所有位都置为0，数组前加一位1
	newDigits := make([]int, len(digits) + 1)
	newDigits[0] = 1
	return newDigits
}