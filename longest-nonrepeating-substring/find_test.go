package substring

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				s: "abcf",
			},
			want: 4,
		},
		{
			name: "ending with max",
			args: args{
				s: "aabcf",
			},
			want: 4,
		},
		{
			name: "given test case 1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "given test case 2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "given test case 3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "length 1",
			args: args{
				s: "A",
			},
			want: 1,
		},
		{
			name: "should be 3",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
