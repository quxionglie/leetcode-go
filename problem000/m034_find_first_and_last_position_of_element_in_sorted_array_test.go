package problem000

import (
	"reflect"
	"sort"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
		{
			name: "",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRange_sort(t *testing.T) {
	leftmost := sort.SearchInts([]int{1, 6, 6, 7}, 6)
	if leftmost != 1 {
		t.Errorf("leftmost() = %v, want %v", leftmost, 1)
	}

	leftmost = sort.SearchInts([]int{1, 2, 6, 6, 7}, 6)
	if leftmost != 2 {
		t.Errorf("leftmost() = %v, want %v", leftmost, 2)
	}

	leftmost = sort.SearchInts([]int{1, 2, 4, 6, 6, 7}, 5)
	if leftmost != 3 {
		t.Errorf("leftmost() = %v, want %v", leftmost, 2)
	}
}
