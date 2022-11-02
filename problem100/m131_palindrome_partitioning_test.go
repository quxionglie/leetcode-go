package problem100

import (
	"reflect"
	"sort"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "",
			args: args{
				s: "aab",
			},
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.args.s)
			var ok bool
			for _, gotStr := range got {
				for _, wantStr := range tt.want {
					if len(gotStr) == len(wantStr) {
						sort.Strings(gotStr)
						sort.Strings(wantStr)
						if reflect.DeepEqual(gotStr, wantStr) {
							ok = true
						}
					}
				}
			}
			if !ok {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
