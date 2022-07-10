package weekgame0710

import (
	"errors"
)

func fillCups(amount []int) int {
	warmNums := amount[0] + 1
	clodNums := amount[1] + 1
	hotNums := amount[2] + 1

	dp := make([][][]int, warmNums)
	for j := 0; j < warmNums; j++ {
		dp[j] = make([][]int, clodNums)
		for k := 0; k < clodNums; k++ {
			dp[j][k] = make([]int, hotNums)
		}
	}
	for i := 0; i < warmNums; i++ {
		dp[i][0][0] = i
	}
	for j := 0; j < clodNums; j++ {
		dp[0][j][0] = j
	}
	for k := 0; k < hotNums; k++ {
		dp[0][0][k] = k
	}

	for i := 0; i < warmNums; i++ {
		for j := 0; j < clodNums; j++ {
			for k := 0; k < hotNums; k++ {
				temp := 3000
				if i > 0 {
					temp = min(temp, dp[i-1][j][k]+1)
				}
				if j > 0 {
					temp = min(temp, dp[i][j-1][k]+1)
				}
				if k > 0 {
					temp = min(temp, dp[i][j][k-1]+1)
				}

				if i > 0 && j > 0 {
					temp = min(temp, dp[i-1][j-1][k]+1)
				}
				if i > 0 && k > 0 {
					temp = min(temp, dp[i-1][j][k-1]+1)
				}
				if j > 0 && k > 0 {
					temp = min(temp, dp[i][j-1][k-1]+1)
				}
				if i > 0 || j > 0 || k > 0 {
					dp[i][j][k] = temp
				}
			}
		}
	}
	return dp[warmNums-1][clodNums-1][hotNums-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type Myheap struct {
	nums []int
}

//创建堆
func NewMyHeap() *Myheap {
	return &Myheap{
		nums: make([]int, 1),
	}
}

func (this *Myheap) Push(val int) {
	//将新添加的结点放置在最后
	this.nums = append(this.nums, val)
	pos := len(this.nums) - 1
	//根据堆的规则，将该结点向上调整
	for pos > 1 && this.nums[pos/2] > this.nums[pos] {
		this.nums[pos/2], this.nums[pos] = this.nums[pos], this.nums[pos/2]
		pos = pos / 2
	}
}

func (this *Myheap) Pop() (int, error) {
	//判断堆中是否还存在结点
	if len(this.nums) <= 1 {
		return 0, errors.New("empty")
	}
	val := this.nums[1]
	//将堆的根结点与最后一个结点交换
	this.nums[1], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[1]
	//删除最后一个结点（原根节点）
	this.nums = this.nums[:len(this.nums)-1]
	//将此时的根结点向下调整到正确的位置
	if len(this.nums) > 1 {
		this.adjustDown(1)
	}
	return val, nil
}

func (this *Myheap) adjustDown(pos int) {
	length := len(this.nums) - 1
	//此处规定下标从1开始，因此下标为0的位置可以用来暂时保存值
	this.nums[0] = this.nums[pos]
	//for循环为找子节点过程
	for i := 2 * pos; i <= length; i *= 2 {
		//找到更小的子节点
		if i < length && this.nums[i] > this.nums[i+1] {
			i++
		}
		//如果存在子节点小于需要调整的结点，则将子节点移动到父节点位置
		if this.nums[0] > this.nums[i] {
			this.nums[pos] = this.nums[i]
			pos = i
		}
	}
	//将结点放置在正确的位置
	this.nums[pos] = this.nums[0]
}

type SmallestInfiniteSet struct {
	heap   *Myheap
	keyMap map[int]bool
}

func Constructor() SmallestInfiniteSet {
	heap := NewMyHeap()
	keyMap := make(map[int]bool, 0)
	nums := 1001
	for i := 1; i < nums; i++ {
		keyMap[i] = true
		heap.Push(i)
	}
	return SmallestInfiniteSet{
		heap:   heap,
		keyMap: keyMap,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	num, _ := this.heap.Pop()
	delete(this.keyMap, num)
	return num
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.keyMap[num]; ok {
		return
	}
	this.keyMap[num] = true
	this.heap.Push(num)
}
