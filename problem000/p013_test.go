package problem000

import (
	"testing"
)

//示例1:
//输入:"III"
//输出: 3

//示例2:
//输入:"IV"
//输出: 4

//示例3:
//输入:"IX"
//输出: 9

//示例4:
//输入:"LVIII"
//输出: 58
//解释: L = 50, V= 5, III = 3.

//示例5:
//输入:"MCMXCIV"
//输出: 1994
//解释: M = 1000, CM = 900, XC = 90, IV = 4.
func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"III"}, 3},
		{"", args{"IV"}, 4},
		{"", args{"IX"}, 9},
		{"", args{"LVIII"}, 58},
		{"", args{"MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
