package weekgame0710

import "testing"

func TestFillCap(t *testing.T) {
	//
	//amount := []int{1, 4, 2}
	//
	//res := fillCups(amount)

	obj := Constructor()
	obj.AddBack(2)
	param_1 := obj.PopSmallest()
	t.Log(param_1)
	param_2 := obj.PopSmallest()
	t.Log(param_2)
	param_3 := obj.PopSmallest()
	t.Log(param_3)
	obj.AddBack(1)
	param_4 := obj.PopSmallest()
	t.Log(param_4)
	param_5 := obj.PopSmallest()
	t.Log(param_5)
	param_6 := obj.PopSmallest()
	t.Log(param_6)
}
