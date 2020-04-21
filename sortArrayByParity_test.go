package leetcode_test

import (
	"fmt"
	"testing"
)

func Test_SortArrayByParity(t *testing.T) {

	tests := map[string]struct {
		input []int
		want  []int
	}{
		"simple": {
			input: []int{3, 1, 2, 4},
			want:  []int{2, 4, 3, 1},
		},
	}

	for name, tc := range tests {
		fmt.Println(name + ": ")
		fmt.Println(tc)
		//out := leetcode.CanJump(tc.input)
		//if out != tc.want {
		//	t.Errorf("name: %s except: %t, actual: %t", name, tc.want, out)
		//}
	}
}
