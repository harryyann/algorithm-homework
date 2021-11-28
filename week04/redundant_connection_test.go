package week04

import "testing"

func TestFindRedundantConnection(t *testing.T) {
	var ins = [][]int{
		{1,2},{2,3},{3,4},{1,4},{1,5},
	}

	redundant := FindRedundantConnection(ins)
	t.Log(redundant)
}
