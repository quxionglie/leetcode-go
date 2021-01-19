package problem100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//-2^31 <= x <= 2^31 - 1
func TestReverse(t *testing.T) {
	//示例 1：
	//输入：x = 123
	//输出：321
	assert.Equal(t, 321, reverse(123))

	//示例 2：
	//输入：x = -123
	//输出：-321
	assert.Equal(t, -321, reverse(-123))

	//示例 3：
	//输入：x = 120
	//输出：21
	assert.Equal(t, 12, reverse(21))

	//示例 4：
	//输入：x = 0
	//输出：0
	assert.Equal(t, 0, reverse(0))

	assert.Equal(t, 0, reverse(2147483647))
	assert.Equal(t, 0, reverse(-2147483648))
}
