package algorithm

import (
	"github.com/stretchr/testify/assert" // 这里引入了 testify
	"testing"
)

func TestBinarySearch(t *testing.T) {
	// 二分查找单测
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "二分查找测试用例",
			args: args{
				nums:   []int{1, 2, 3, 4},
				target: 3,
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got := BinarySearch(tt.args.nums, tt.args.target)
		// 改写使用断言的方式
		assert.Equal(t, tt.want, got)
	}
}
