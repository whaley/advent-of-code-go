package _019

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
		name   string
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}, "1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2)."},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}, "2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6)"},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}, "2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801)."},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, "1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99."},
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, "Default case"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, RunIntCodeProgram(test.input), test.name)
		})
	}
}
