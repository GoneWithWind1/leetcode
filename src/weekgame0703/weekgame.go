package weekgame0703

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func spiralMatrix(m int, n int, head *ListNode) [][]int {

	nums := make([][]int, m)
	for i := range nums {
		nums[i] = make([]int, n)
	}
	roundTime := (min(m, n) + 1) / 2
	p := head
	for i := 0; i < roundTime; i++ {
		if n-1-i >= i {
			p = fillMatrixHorizontal(nums, i, i, n-1-i, p, 0)
		}
		if m-1-i >= i+1 {
			p = fillMatrixLongitudinal(nums, n-1-i, i+1, m-1-i, p, 0)
		}
		if m-1-i > i && n-2-i >= i {
			p = fillMatrixHorizontal(nums, m-1-i, n-2-i, i, p, 1)
		}
		if n-1-i > i && m-2-i >= i+1 {
			p = fillMatrixLongitudinal(nums, i, m-2-i, i+1, p, 1)
		}
	}
	return nums
}

// falg 表示填充正反，0表示正向填充，1表示反向填充
func fillMatrixHorizontal(nums [][]int, row int, start int, end int, p *ListNode, flag int) *ListNode {
	if flag == 0 {
		for col := start; col <= end; col++ {
			if p != nil {
				nums[row][col] = p.Val
				p = p.Next
			} else {
				nums[row][col] = -1
			}
		}

	} else {
		for col := start; col >= end; col-- {
			if p != nil {
				nums[row][col] = p.Val
				p = p.Next
			} else {
				nums[row][col] = -1
			}
		}
	}
	return p
}

// falg 表示填充正反，0表示正向填充，1表示反向填充
func fillMatrixLongitudinal(nums [][]int, col int, start int, end int, p *ListNode, flag int) *ListNode {
	if flag == 0 {
		for row := start; row <= end; row++ {
			if p != nil {
				nums[row][col] = p.Val
				p = p.Next
			} else {
				nums[row][col] = -1
			}
		}

	} else {
		for row := start; row >= end; row-- {
			if p != nil {
				nums[row][col] = p.Val
				p = p.Next
			} else {
				nums[row][col] = -1
			}
		}
	}
	return p
}
