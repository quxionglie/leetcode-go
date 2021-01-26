package problem100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	//示例1:
	//输入:"III"
	//输出: 3
	assert.Equal(t, 3, romanToInt("III"))

	//示例2:
	//输入:"IV"
	//输出: 4
	assert.Equal(t, 4, romanToInt("IV"))

	//示例3:
	//输入:"IX"
	//输出: 9
	assert.Equal(t, 9, romanToInt("IX"))

	//示例4:
	//输入:"LVIII"
	//输出: 58
	//解释: L = 50, V= 5, III = 3.
	assert.Equal(t, 58, romanToInt("LVIII"))

	//示例5:
	//输入:"MCMXCIV"
	//输出: 1994
	//解释: M = 1000, CM = 900, XC = 90, IV = 4.
	assert.Equal(t, 1994, romanToInt("MCMXCIV"))
}
