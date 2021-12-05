package aoc

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

type dynamicGrid struct {
	xMax, yMax int
	grid       [][]int
	overlaps   map[string]bool
}

func (g *dynamicGrid) growX(newMax int) {
	newMax += 1
	if newMax <= g.xMax {
		return
	}

	lengthToAdd := newMax - g.xMax

	for i := range g.grid {
		toAdd := make([]int, lengthToAdd)
		g.grid[i] = append(g.grid[i], toAdd...)
	}

	g.xMax = newMax
}

func (g *dynamicGrid) growY(newMax int) {
	newMax += 1
	if newMax <= g.yMax {
		return
	}

	rowsToAdd := newMax - g.yMax

	for i := 0; i < rowsToAdd; i++ {
		newRow := make([]int, g.xMax)
		g.grid = append(g.grid, newRow)
	}

	g.yMax = newMax
}

func (g *dynamicGrid) markPoints(x1, y1, x2, y2 int) {
	g.growY(y1)
	g.growY(y2)
	g.growX(x1)
	g.growX(x2)

	xDelta := x2 - x1
	yDelta := y2 - y1

	xDir := 0
	if xDelta > 0 {
		xDir = 1
	} else if xDelta < 0 {
		xDir = -1
	}
	yDir := 0
	if yDelta > 0 {
		yDir = 1
	} else if yDelta < 0 {
		yDir = -1
	}

	xMarksToAdd, yMarksToAdd := abs(xDelta), abs(yDelta)
	xMarks, yMarks := 0, 0
	xx, yy := x1, y1
	for {
		g.grid[yy][xx]++
		if g.grid[yy][xx] > 1 {
			g.overlaps[fmt.Sprintf("%d,%d", xx, yy)] = true
		}

		if xMarks < xMarksToAdd {
			xx += xDir
		}
		if yMarks < yMarksToAdd {
			yy += yDir
		}

		xMarks++
		yMarks++

		if xMarks > xMarksToAdd && yMarks > yMarksToAdd {
			break
		}
	}
}

func (g *dynamicGrid) getNumOverlaps() int {
	return len(g.overlaps)
}

func hydrothermalVenturePartOne(vents []string) int {
	grid := dynamicGrid{
		grid:     [][]int{},
		overlaps: map[string]bool{},
	}

	for _, vent := range vents {
		split := strings.Split(vent, " -> ")
		coords := make([]int, 0, 4)

		for _, s := range split {
			innerSplit := strings.Split(s, ",")
			for _, z := range innerSplit {
				n, err := strconv.Atoi(z)
				if err != nil {
					panic("unable to convert coord to int")
				}
				coords = append(coords, n)
			}
		}

		x1, y1 := coords[0], coords[1]
		x2, y2 := coords[2], coords[3]

		// For now, only consider horizontal and vertical lines: lines where either x1 = x2
		// or y1 = y2.
		if x1 != x2 && y1 != y2 {
			continue
		}

		grid.markPoints(x1, y1, x2, y2)
	}

	return grid.getNumOverlaps()
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func TestDayFive(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/05.txt")
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
					"0,9 -> 5,9",
					"8,0 -> 0,8",
					"9,4 -> 3,4",
					"2,2 -> 2,1",
					"7,0 -> 7,4",
					"6,4 -> 2,0",
					"0,9 -> 2,9",
					"3,4 -> 1,4",
					"0,0 -> 8,8",
					"5,5 -> 8,2",
				},
				expected: 5,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, hydrothermalVenturePartOne)
	})
}
