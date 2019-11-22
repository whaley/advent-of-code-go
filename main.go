package main

import (
	"fmt"
	"github.com/whaley/advent-of-code/common"
	"strings"
)

func main() {
	day01()
	day02()
	day03()
	//day04()
}

func day01() {
	freqs := common.DelimitByNewLine(common.ReadFully("day01.txt"))
	for idx, freq := range freqs {
		freqs[idx] = strings.TrimSpace(freq)
	}

	pt1Answer := computeFrequency(freqs)
	pt2Answer := findRepeatingFrequency(freqs)

	fmt.Printf("Day 01 : Part 01  Answer:\n\t%d\n", pt1Answer)
	fmt.Printf("Day 01 : Part 02  Answer:\n\t%d\n", pt2Answer)
}

func day02() {
	lines := common.DelimitByNewLine(common.ReadFully("day02.txt"))

	pt1Answer := computeChecksum(lines)
	pt2Answer := day02Pt2(lines)

	fmt.Printf("Day 02 : Part 01  Answer:\n\t%d\n", pt1Answer)
	fmt.Printf("Day 02 : Part 02  Answer:\n\t%s\n", pt2Answer)
}

func day03() {
	var rects []Rectangle
	for _, line := range common.DelimitByNewLine(common.ReadFully("day03.txt")) {
		rect, err := createRect(line)
		common.PanicOnError(err)
		rects = append(rects, rect)
	}
	grid := NewGrid(rects)
	day0301Points := grid.FilterBy(func(claims map[int]bool) bool {
		return len(claims) > 1
	})

	id, err := FindNonOverlappingClaim(rects)
	common.PanicOnError(err)
	fmt.Printf("Day 03 : Part 01  Answer:\n\t%d\n", len(day0301Points))
	fmt.Printf("Day 03 : Part 02  Answer:\n\t%d\n", id)
}

func day04() {
	lines := common.DelimitByNewLine(common.ReadFully("day04.txt"))
	err := SolveDay04Pt1(lines)
	if err != nil {
		panic(err)
	}
}
