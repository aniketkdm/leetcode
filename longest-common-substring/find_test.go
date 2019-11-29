package substring

import "testing"

func Test_longestCommonSubstring(t *testing.T) {
	type args struct {
		s string
		r string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				s: "abcbd",
				r: "dbcba",
			},
			want: "bcb",
		},
		{
			name: "different lengths",
			args: args{
				s: "abcbd",
				r: "bcb",
			},
			want: "bcb",
		},
		{
			name: "not palindrome",
			args: args{
				s: "abcxd",
				r: "bcx",
			},
			want: "bcx",
		},
		{
			name: "one blank",
			args: args{
				s: "abcxd",
				r: "",
			},
			want: "",
		},
		{
			name: "both blank",
			args: args{
				s: "abcxd",
				r: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonSubstring(tt.args.s, tt.args.r); got != tt.want {
				t.Errorf("longestCommonSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
