package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/MEX7/leetcode"
)

func Test_GameOfLife(t *testing.T) {

	tests := map[string]struct {
		input [][]int
		want  [][]int
	}{
		"simple": {
			input: [][]int{[]int{0, 1, 0}, []int{0, 0, 1}, []int{1, 1, 1}, []int{0, 0, 0}},
			want:  [][]int{[]int{0, 0, 0}, []int{1, 0, 1}, []int{0, 1, 1}, []int{0, 1, 0}},
		},
	}

	for name, tc := range tests {
		fmt.Println(name + ": ")
		out := leetcode.GameOfLife(tc.input)
		for x, _ := range out {
			for y, _ := range out[x] {
				if out[x][y] != tc.want[x][y] {
					t.Errorf("x: %d, y: %d, except: %d, actual: %d", x, y, tc.want[x][y], out[x][y])
				}
			}
		}
	}
}
