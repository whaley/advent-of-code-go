package _019

import (
	"fmt"
	"log"
)

const ADD = 1
const MULT = 2
const HALT = 99

func RunIntCodeProgram(intCode []int) []int {
	wc := make([]int, len(intCode))
	copy(wc, intCode)

	opIdx := 0
	opCode := wc[opIdx]

	for {
		switch opCode {
		case ADD:
			wc[wc[opIdx+3]] = wc[wc[opIdx+1]] + wc[wc[opIdx+2]]
		case MULT:
			wc[wc[opIdx+3]] = wc[wc[opIdx+1]] * wc[wc[opIdx+2]]
		case HALT:
			return wc
		default:
			log.Panicf("Received an opcode that not one of %v, %v, or %v", ADD, MULT, HALT)
		}

		opIdx += 4
		if opIdx < len(wc) {
			opCode = wc[opIdx]
		} else {
			break
		}
	}

	return wc
}

func Run2019Day02() {
	intCode := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 9, 19, 23, 1, 13, 23, 27, 1, 5, 27, 31, 2, 31, 6, 35, 1, 35, 5, 39, 1, 9, 39, 43, 1, 43, 5, 47, 1, 47, 5, 51, 2, 10, 51, 55, 1, 5, 55, 59, 1, 59, 5, 63, 2, 63, 9, 67, 1, 67, 5, 71, 2, 9, 71, 75, 1, 75, 5, 79, 1, 10, 79, 83, 1, 83, 10, 87, 1, 10, 87, 91, 1, 6, 91, 95, 2, 95, 6, 99, 2, 99, 9, 103, 1, 103, 6, 107, 1, 13, 107, 111, 1, 13, 111, 115, 2, 115, 9, 119, 1, 119, 6, 123, 2, 9, 123, 127, 1, 127, 5, 131, 1, 131, 5, 135, 1, 135, 5, 139, 2, 10, 139, 143, 2, 143, 10, 147, 1, 147, 5, 151, 1, 151, 2, 155, 1, 155, 13, 0, 99, 2, 14, 0, 0}
	//To do this, before running the program, replace position 1 with the value 12 and replace position 2 with the value 2.
	intCode[1] = 12
	intCode[2] = 2

	output := RunIntCodeProgram(intCode)

	fmt.Printf("Day 02 : Part 01  Answer:\n\t%d\n", output[0]) //What value is left at position 0 after the program halts?
	fmt.Printf("Day 02 : Part 02  Answer:\n\t%d\n", "TODO") //TODO: do part 2
}
