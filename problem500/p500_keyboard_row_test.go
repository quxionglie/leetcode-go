package problem500

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "",
			args: args{
				words: []string{"Hello", "Alaska", "Dad", "Peace"},
			},
			wantAns: []string{"Alaska", "Dad"},
		},
		{
			name: "",
			args: args{
				words: []string{"amk"},
			},
			wantAns: []string{},
		},
		{
			name: "",
			args: args{
				words: []string{"adsdf", "sfd"},
			},
			wantAns: []string{"adsdf", "sfd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findWords(tt.args.words); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findWords() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
