package aoc

import (
	"fmt"
	"testing"
	"time"
)

type smokeGrid struct {
	values        []rune
	width, height int
}

func (g *smokeGrid) setValue(x, y int, value rune) {
	g.values[y*g.width+x] = value
}

func (g *smokeGrid) getValue(x, y int) rune {
	return g.values[y*g.width+x]
}

func (g *smokeGrid) getLowestPoints() []rune {
	lowPoints := []rune{}

	for y := 0; y < g.height; y++ {
	adjLoop:
		for x := 0; x < g.width; x++ {
			value := g.getValue(x, y)
			adjacent := g.getAdjacent(x, y)

			for _, adj := range adjacent {
				if adj <= value {
					continue adjLoop
				}
			}

			lowPoints = append(lowPoints, value)
		}
	}

	return lowPoints
}

func (g *smokeGrid) getAdjacent(x, y int) []rune {
	adjacent := []rune{}

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

func smokeBasinPartOne(heights []string) int {
	gridValues := make([]rune, len(heights)*len(heights[0]))
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
		total += int(point - '0')
		total += 1
	}

	return total
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
}
