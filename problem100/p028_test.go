package problem100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
示例 1：
输入：haystack = "hello", needle = "ll"
输出：2

示例 2：
输入：haystack = "aaaaa", needle = "bba"
输出：-1

示例 3：
输入：haystack = "", needle = ""
输出：0
*/
func TestStrStr(t *testing.T) {
	assert.Equal(t, 2, strStr("hello", "ll"))
	assert.Equal(t, -1, strStr("aaaaa", "bba"))
	assert.Equal(t, 0, strStr("", ""))
	assert.Equal(t, -1, strStr("aaa", "aaaa"))
	assert.Equal(t, -1, strStr("mississippi", "issipi"))
}
