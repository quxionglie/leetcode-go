package problem500

import "testing"

func Test_checkRecord(t *testing.T) {
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
			args: args{"PPALLP"},
			want: true,
		},
		{
			name: "",
			args: args{"PPALLL"},
			want: false,
		},
		{
			name: "",
			args: args{"LLL"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRecord(tt.args.s); got != tt.want {
				t.Errorf("checkRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
