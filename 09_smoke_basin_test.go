package aoc

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

type smokeGridPoint struct {
	x, y  int
	value rune
}

type smokeGrid struct {
	values        []smokeGridPoint
	width, height int
}

func (g *smokeGrid) setValue(x, y int, value rune) {
	g.values[y*g.width+x] = smokeGridPoint{
		x:     x,
		y:     y,
		value: value,
	}
}

func (g *smokeGrid) getValue(x, y int) smokeGridPoint {
	return g.values[y*g.width+x]
}

func (g *smokeGrid) getLowestPoints() []smokeGridPoint {
	lowPoints := []smokeGridPoint{}

	for y := 0; y < g.height; y++ {
	adjLoop:
		for x := 0; x < g.width; x++ {
			point := g.getValue(x, y)
			adjacent := g.getAdjacent(x, y)

			for _, adj := range adjacent {
				if adj.value <= point.value {
					continue adjLoop
				}
			}

			lowPoints = append(lowPoints, point)
		}
	}

	return lowPoints
}

func (g *smokeGrid) getAdjacent(x, y int) []smokeGridPoint {
	adjacent := []smokeGridPoint{}

	if x-1 >= 0 {
		adjacent = append(adjacent, g.getValue(x-1, y))
	}
	if y-1 >= 0 {
		adjacent = append(adjacent, g.getValue(x, y-1))
	}
	if x+1 < g.width {
		adjacent = append(adjacent, g.getValue(x+1, y))
	}
	if y+1 < g.height {
		adjacent = append(adjacent, g.getValue(x, y+1))
	}

	return adjacent
}

func (g *smokeGrid) getBasinSizes(lowPoints []smokeGridPoint) []int {
	basinSizes := []int{}

	for _, point := range lowPoints {
		adjacent := g.getAdjacent(point.x, point.y)

		seen := map[string]bool{}
		seen[fmt.Sprintf("%d,%d", point.x, point.y)] = true

		i := 0
		for i < len(adjacent) {
			innerPoint := adjacent[i]

			_, ok := seen[fmt.Sprintf("%d,%d", innerPoint.x, innerPoint.y)]

			if innerPoint.value-'0' != 9 && !ok {
				seen[fmt.Sprintf("%d,%d", innerPoint.x, innerPoint.y)] = true
				innerAdjacent := g.getAdjacent(innerPoint.x, innerPoint.y)
				adjacent = append(adjacent, innerAdjacent...)
			}

			i++
		}

		basinSizes = append(basinSizes, len(seen))
	}

	return basinSizes
}

func smokeBasinPartOne(heights []string) int {
	gridValues := make([]smokeGridPoint, len(heights)*len(heights[0]))
	grid := smokeGrid{
		values: gridValues,
		width:  len(heights[0]),
		height: len(gridValues) / len(heights[0]),
	}

	for y, row := range heights {
		for x, char := range row {
			grid.setValue(x, y, char)
		}
	}

	lowPoints := grid.getLowestPoints()

	total := 0
	for _, point := range lowPoints {
		total += int(point.value - '0')
		total += 1
	}

	return total
}

func smokeBasinPartTwo(heights []string) int {
	gridValues := make([]smokeGridPoint, len(heights)*len(heights[0]))
	grid := smokeGrid{
		values: gridValues,
		width:  len(heights[0]),
		height: len(gridValues) / len(heights[0]),
	}

	for y, row := range heights {
		for x, char := range row {
			grid.setValue(x, y, char)
		}
	}

	lowPoints := grid.getLowestPoints()
	basins := grid.getBasinSizes(lowPoints)
	sort.Ints(basins)

	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func TestDayNine(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/09.txt")
	if err != nil {
		t.Fatalf("failed to read input")
	}

	runTests := func(t *testing.T, tests map[string]testConfig, fn func([]string) int) {
		for name, cfg := range tests {
			cfg := cfg
			t.Run(name, func(t *testing.T) {
				start := time.Now()
				output := fn(cfg.input)
				finish := time.Since(start)
				if cfg.logResult {
					t.Log(fmt.Sprintf("\nsolution:\t%v\nelapsed time:\t%s", output, finish))
					return
				}

				if output != cfg.expected {
					t.Fatalf("Incorrect output - got: %v, want: %v", output, cfg.expected)
				}
			})
		}
	}

	t.Run("part one", func(t *testing.T) {
		tests := map[string]testConfig{
			"test 1": {
				input: []string{
					"2199943210",
					"3987894921",
					"9856789892",
					"8767896789",
					"9899965678",
				},
				expected: 15,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, smokeBasinPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test 1": {
				input: []string{
					"2199943210",
					"3987894921",
					"9856789892",
					"8767896789",
					"9899965678",
				},
				expected: 1134,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, smokeBasinPartTwo)
	})
}
