package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/MEX7/leetcode"
)

func Test_CanJump(t *testing.T) {

	tests := map[string]struct {
		input []int
		want  bool
	}{
		"v1": {
			input: []int{2, 3, 1, 1, 4},
			want:  true,
		},
		"v2": {
			input: []int{3, 2, 1, 0, 4},
			want:  false,
		},
		"v3": {
			input: []int{0, 2, 1, 1, 4},
			want:  false,
		},
		"v4": {
			input: []int{2, 0, 0},
			want:  true,
		},
		"v5": {
			input: []int{2, 5, 0, 0},
			want:  true,
		},
	}

	for name, tc := range tests {
		fmt.Println(name + ": ")
		out := leetcode.CanJump(tc.input)
		if out != tc.want {
			t.Errorf("name: %s except: %t, actual: %t", name, tc.want, out)
		}
	}
}
