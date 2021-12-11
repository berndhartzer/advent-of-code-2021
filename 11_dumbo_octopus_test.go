package aoc

import (
	"fmt"
	"testing"
	"time"
)

type octoPoint struct {
	x, y, value int
	lastFlash   int
}

func (p *octoPoint) increase(step int) bool {
	if p.lastFlash != step {
		p.value += 1

		if p.value > 9 {
			p.value = 0
			p.lastFlash = step
			return true
		}
	}

	return false
}

type octoGrid struct {
	values        []*octoPoint
	width, height int
}

func (g *octoGrid) setValue(x, y int, value rune) {
	g.values[y*g.width+x] = &octoPoint{
		x:         x,
		y:         y,
		value:     int(value - '0'),
		lastFlash: -1,
	}
}

func (g *octoGrid) getValue(x, y int) *octoPoint {
	return g.values[y*g.width+x]
}

func (g *octoGrid) step(num int) int {
	adjacent := []*octoPoint{}
	flashes := 0

	for _, point := range g.values {
		startedFlashing := point.increase(num)
		if startedFlashing {
			flashes++
			adjacent = append(adjacent, g.getAdjacent(point.x, point.y)...)
		}
	}

	for len(adjacent) > 0 {
		point := adjacent[0]
		adjacent = adjacent[1:]

		startedFlashing := point.increase(num)
		if startedFlashing {
			flashes++
			adjacent = append(adjacent, g.getAdjacent(point.x, point.y)...)
		}
	}

	return flashes
}

func (g *octoGrid) getAdjacent(x, y int) []*octoPoint {
	adjacent := []*octoPoint{}

	if x-1 >= 0 {
		adjacent = append(adjacent, g.getValue(x-1, y))
		if y-1 >= 0 {
			adjacent = append(adjacent, g.getValue(x-1, y-1))
		}
	}
	if y-1 >= 0 {
		adjacent = append(adjacent, g.getValue(x, y-1))
		if x+1 < g.width {
			adjacent = append(adjacent, g.getValue(x+1, y-1))
		}
	}
	if x+1 < g.width {
		adjacent = append(adjacent, g.getValue(x+1, y))
		if y+1 < g.height {
			adjacent = append(adjacent, g.getValue(x+1, y+1))
		}
	}
	if y+1 < g.height {
		adjacent = append(adjacent, g.getValue(x, y+1))
		if x-1 >= 0 {
			adjacent = append(adjacent, g.getValue(x-1, y+1))
		}
	}

	return adjacent
}

func dumboOctopusPartOne(lines []string) int {
	grid := octoGrid{
		values: make([]*octoPoint, 100),
		width:  10,
		height: 10,
	}

	for y, row := range lines {
		for x, char := range row {
			grid.setValue(x, y, char)
		}
	}

	steps := 100
	total := 0

	for i := 0; i < steps; i++ {
		total += grid.step(i)
	}

	return total
}

func TestDayEleven(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/11.txt")
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
					"5483143223",
					"2745854711",
					"5264556173",
					"6141336146",
					"6357385478",
					"4167524645",
					"2176841721",
					"6882881134",
					"4846848554",
					"5283751526",
				},
				expected: 1656,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, dumboOctopusPartOne)
	})
}
