package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/MEX7/leetcode"
)

func Test_Trap(t *testing.T) {

	tests := map[string]struct {
		input []int
		want  int
	}{
		"v1": {
			input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:  6,
		},
		"v2": {
			input: []int{0, 1, 0, 2, 0, 0, 0, 3, 2, 1, 2, 1},
			want:  8,
		},
		"v3": {
			input: []int{4, 2, 3},
			want:  1,
		},
		"v4": {
			input: []int{9, 8, 2, 6},
			want:  4,
		},
	}

	for name, tc := range tests {
		fmt.Println(name + ": ")
		out := leetcode.Trap(tc.input)
		if out != tc.want {
			t.Errorf("name: %s except: %d, actual: %d", name, tc.want, out)
		}
	}
}
