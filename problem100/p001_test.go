package problem100

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
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
func TestTwoSum(t *testing.T) {
	var out []int
	out = twoSum([]int{2, 7, 11, 15}, 9)
	assert.True(t, array_sorted_equal([]int{0, 1}, out))

	out = twoSum([]int{3, 2, 4}, 6)
	assert.True(t, array_sorted_equal([]int{1, 2}, out))

	out = twoSum([]int{3, 3}, 6)
	assert.True(t, array_sorted_equal([]int{0, 1}, out))

	out = twoSum([]int{1, 3}, 6)
	assert.True(t, array_sorted_equal(nil, out))
	assert.True(t, array_sorted_equal([]int{1, 2}, []int{2, 1}))
}

func array_sorted_equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	a_copy := make([]int, len(a))
	b_copy := make([]int, len(b))

	copy(a_copy, a)
	copy(b_copy, b)

	sort.Ints(a_copy)
	sort.Ints(b_copy)

	return reflect.DeepEqual(a_copy, b_copy)
}
