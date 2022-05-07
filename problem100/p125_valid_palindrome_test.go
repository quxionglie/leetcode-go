package problem100

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	str := "Go爱好者"
	fmt.Printf("11=%d, \n", len([]rune(str)))
	fmt.Printf("22=%d, \n", len(str))

	//for _, c := range str {
	//	fmt.Printf("%s,", string(c))
	//	fmt.Printf("%q,", []rune(str)[2])
	//	//fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	//}

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{"A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "",
			args: args{"race a car"},
			want: false,
		},
		{
			name: "",
			args: args{""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
