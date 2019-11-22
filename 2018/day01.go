package _018

import (
	"fmt"
	"github.com/whaley/advent-of-code/common"
	"log"
	"strconv"
	"strings"
)

func computeFrequency(freqChanges []string) int {
	freq := 0
	for _, changeString := range freqChanges {
		change := stringToInt(changeString)
		freq += change
	}
	return freq
}

func findRepeatingFrequency(freqChanges []string) int {
	freq := 0
	seen := make(map[int]bool)
	idx := 0

	for !seen[freq] {
		seen[freq] = true

		change := stringToInt(freqChanges[idx])
		freq += change

		idx++
		if idx == len(freqChanges) {
			idx = 0
		}
	}
	return freq
}

func stringToInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func RunDay01() {
	freqs := common.DelimitByNewLine(common.ReadFully("static/2018/day01.txt"))
	for idx, freq := range freqs {
		freqs[idx] = strings.TrimSpace(freq)
	}

	pt1Answer := computeFrequency(freqs)
	pt2Answer := findRepeatingFrequency(freqs)

	fmt.Printf("Day 01 : Part 01  Answer:\n\t%d\n", pt1Answer)
	fmt.Printf("Day 01 : Part 02  Answer:\n\t%d\n", pt2Answer)
}

