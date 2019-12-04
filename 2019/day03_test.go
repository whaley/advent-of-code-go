package _019

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDistanceToCrossClosestToPort(t *testing.T) {
	tests := []struct {
		lines                      []string
		expectedDistance           int
		expectedLeastCombinedSteps int
		name                       string
	}{
		{[]string{"R8,U5,L5,D3", "U7,R6,D4,L4"}, 6, 30, "Simple case"},
		{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 159, 610, "First advanced case"},
		{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 135, 410, "Second advanced case"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			leastDistance, leastCombinedSteps := FindDistanceToCrossClosestToPort(test.lines)
			assert.Equal(t, test.expectedDistance, leastDistance, test.name)
			assert.Equal(t, test.expectedLeastCombinedSteps, leastCombinedSteps, test.name)
		})
	}
}

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		pointA   Coord
		pointB   Coord
		distance int
		name     string
	}{
		{Coord{0, 0}, Coord{0, 0}, 0, "Same Point"},
		{Coord{0, 5}, Coord{0, 0}, 5, "Y only 10"},
		{Coord{0, -5}, Coord{0, 0}, 5, "Y only -5"},
		{Coord{5, 5}, Coord{0, 0}, 10, "X and Y 5"},
		{Coord{5, -5}, Coord{0, 0}, 10, "X and Y 5"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.distance, ManhattanDistance(test.pointA, test.pointB), test.name)
		})
	}
}
