package _019

import (
	"fmt"
	"github.com/whaley/advent-of-code/common"
	"math"
	"strconv"
	"strings"
)

func CalcFuelForMass(mass int) int {
	x := float64(mass) / 3
	y := int(math.Floor(x))
	return y - 2
}

func recur(mass int, accum int) int {
	fuelNeeded := CalcFuelForMass(mass)
	if fuelNeeded <= 0 {
		return accum
	} else {
		return recur(fuelNeeded, accum+fuelNeeded)
	}
}

func CalcTotalFuelRequiredForMass(mass int) int {
	return recur(mass, 0)
}

func CalcFuelRequiredForAllMasses(masses []int, f func(int) int) int {
	fuel := 0
	for _, mass := range masses {
		fuel += f(mass)
	}
	return fuel
}

func Run2019Day01() {
	lines := common.DelimitByNewLine(common.ReadFully("static/2019/day01.txt"))
	masses := make([]int, len(lines))
	for i, line := range lines {
		mass, err := strconv.Atoi(strings.TrimSpace(line))
		common.PanicOnError(err)
		masses[i] = mass
	}

	fmt.Printf("Day 01 : Part 01  Answer:\n\t%d\n", CalcFuelRequiredForAllMasses(masses, CalcFuelForMass))
	fmt.Printf("Day 01 : Part 02  Answer:\n\t%d\n", CalcFuelRequiredForAllMasses(masses, CalcTotalFuelRequiredForMass))
}
