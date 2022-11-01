package problem100

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "",
			args: args{
				board: [][]byte{
					{'X'},
				},
			},
			want: [][]byte{
				{'X'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			if !reflect.DeepEqual(tt.want, tt.args.board) {
				t.Errorf("solve() = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}
