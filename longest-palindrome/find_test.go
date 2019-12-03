package palindrome

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
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
			},
			want: "bcb",
		},
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "no palindrome",
			args: args{
				s: "abcdb",
			},
			want: "b",
		},
		{
			name: "different from longest substring",
			args: args{
				s: "aaczecaa",
			},
			want: "aa",
		},
		{
			name: "different from longest substring and greater than palindrome in substring",
			args: args{
				s: "aaczecaabcb",
			},
			want: "bcb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				s: "abcd",
			},
			want: "dcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s     string
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "basic",
			args: args{
				s:     "bcb",
				start: 0,
				end:   2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
