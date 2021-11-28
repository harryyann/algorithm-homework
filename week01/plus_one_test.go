package week01

import "testing"

func TestPlusOne(t *testing.T) {
	digits := []int{9,9,9}
	r := plusOne(digits)
	t.Log(r)
}