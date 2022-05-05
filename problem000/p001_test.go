package problem000

import (
	"reflect"
	"testing"
)

//示例 1：
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

//示例 2：
//输入：nums = [3,2,4], target = 6
//输出：[1,2]

//示例 3：
//输入：nums = [3,3], target = 6
//输出：[0,1]
//
//提示：
//2 <= nums.length <= 103
//-109 <= nums[i] <= 109
//-109 <= target <= 109
//只会存在一个有效答案
func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"01", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"02", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"03", args{[]int{3, 3}, 6}, []int{0, 1}},
		{"04", args{[]int{1, 3}, 6}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
