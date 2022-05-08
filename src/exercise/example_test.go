package exercise

import "testing"

func TestExample(t *testing.T) {

	nums := []int{1, 2, 3, 3}
	res := FindDuplicates(nums)
	t.Log(res)
}
