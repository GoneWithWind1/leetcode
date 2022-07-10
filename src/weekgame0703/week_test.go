package weekgame0703

import "testing"

func TestSpiralMatrix(t *testing.T) {

	nodeList := []int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0}
	head := &ListNode{}
	p := head
	for _, val := range nodeList {
		node := &ListNode{
			Val: val,
		}
		p.Next = node
		p = node
	}

	res := spiralMatrix(3, 5, head.Next)
	t.Log(res)
}
