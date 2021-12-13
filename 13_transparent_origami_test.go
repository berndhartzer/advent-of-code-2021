package aoc

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

type origamiGrid struct {
	xMax, yMax int
	values     [][]int
}

func (g *origamiGrid) markPoint(x, y int) {
	g.grow(x, y)
	g.values[y][x]++
}

func (g *origamiGrid) grow(x, y int) {
	newX := x + 1
	newY := y + 1

	if newY > len(g.values) {
		// add more rows
		yToAdd := newY - g.yMax

		for i := 0; i < yToAdd; i++ {
			newRow := make([]int, g.xMax)
			g.values = append(g.values, newRow)
		}

		g.yMax = newY
	}

	if newX > len(g.values[0]) {
		// add more columns
		xToAdd := newX - g.xMax
		for i := range g.values {
			toAdd := make([]int, xToAdd)
			g.values[i] = append(g.values[i], toAdd...)
		}

		g.xMax = newX
	}
}

func (g *origamiGrid) fold(instructions string) {
	split := strings.Split(instructions, "=")
	axis := string(split[0][len(split[0])-1])

	line, err := strconv.Atoi(split[1])
	if err != nil {
		panic("converting fold to int")
	}

	switch axis {
	case "x":
		g.foldX(line)
	case "y":
		g.foldY(line)
	}
}

func (g *origamiGrid) foldY(line int) {
	bottomFold := g.values[line+1:]
	foldLine := (len(g.values) - (line + 1)) - 1

	for i, row := range bottomFold {
		rowToUpdate := foldLine - i
		for j := 0; j < len(row); j++ {
			g.values[rowToUpdate][j] = g.values[rowToUpdate][j] + row[j]
		}
	}

	g.values = g.values[:len(g.values)-(line+1)]
}

func (g *origamiGrid) foldX(line int) {
	for i := 0; i < len(g.values); i++ {
		toFold := g.values[i][line+1:]
		foldLine := (len(g.values[i]) - (line + 1)) - 1

		for j, point := range toFold {
			g.values[i][foldLine-j] = g.values[i][foldLine-j] + point
		}

		g.values[i] = g.values[i][:len(g.values[i])-(line+1)]
	}
}

func (g *origamiGrid) visiblePoints() int {
	total := 0
	for _, row := range g.values {
		for _, point := range row {
			if point > 0 {
				total++
			}
		}
	}
	return total
}

func (g *origamiGrid) print() {
	for _, row := range g.values {
		rowStr := ""
		for _, point := range row {
			if point > 0 {
				rowStr += "*"
			} else {
				rowStr += "_"
			}
		}
		fmt.Println(rowStr)
	}
}

func transparentOrigamiPartOne(input []string) int {
	grid := origamiGrid{}

	i := 0
	for ; i < len(input); i++ {
		line := input[i]
		if line == "" {
			break
		}
		split := strings.Split(line, ",")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			panic("converting x to int")
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			panic("converting y to int")
		}

		grid.markPoint(x, y)
	}
	i++

	grid.fold(input[i])

	return grid.visiblePoints()
}

func transparentOrigamiPartTwo(input []string) int {
	grid := origamiGrid{}

	i := 0
	for ; i < len(input); i++ {
		line := input[i]
		if line == "" {
			break
		}
		split := strings.Split(line, ",")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			panic("converting x to int")
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			panic("converting y to int")
		}

		grid.markPoint(x, y)
	}
	i++

	for ; i < len(input); i++ {
		grid.fold(input[i])
	}

	grid.print()

	return 0
}

func TestDayThirteen(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/13.txt")
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
					"6,10",
					"0,14",
					"9,10",
					"0,3",
					"10,4",
					"4,11",
					"6,0",
					"6,12",
					"4,1",
					"0,13",
					"10,12",
					"3,4",
					"3,0",
					"8,4",
					"1,10",
					"2,14",
					"8,10",
					"9,0",
					"",
					"fold along y=7",
					"fold along x=5",
				},
				expected: 17,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, transparentOrigamiPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, transparentOrigamiPartTwo)
	})
}
