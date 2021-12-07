package aoc

import (
	"fmt"
	"testing"
	"time"
)

func theTreacheryOfWhalesPartOne(crabs []int) int {
	positions := []int{}

	for _, crab := range crabs {
		if crab+1 > len(positions) {
			grow := crab + 1 - len(positions)
			additions := make([]int, grow)
			positions = append(positions, additions...)
		}

		positions[crab]++
	}

	fuelToPosition := make([]int, len(positions))
	for pos := range fuelToPosition {
		totalFuelToPos := 0
		for crab, count := range positions {
			diff := abs(crab - pos)
			totalFuelToPos += diff * count
		}

		fuelToPosition[pos] = totalFuelToPos
	}

	least := -1
	for _, fuel := range fuelToPosition {
		if least == -1 {
			least = fuel
		}

		if fuel < least {
			least = fuel
		}
	}

	return least
}

func theTreacheryOfWhalesPartTwo(crabs []int) int {
	positions := []int{}

	for _, crab := range crabs {
		if crab+1 > len(positions) {
			grow := crab + 1 - len(positions)
			additions := make([]int, grow)
			positions = append(positions, additions...)
		}

		positions[crab]++
	}

	fuelToPosition := make([]int, len(positions))

	// Improves performance by about 10x, from ~750ms to ~75ms
	fuelToMoveOneCache := map[int]int{}

	for pos := range fuelToPosition {
		totalFuelToPos := 0
		for crab, count := range positions {
			diff := abs(crab - pos)

			fuelToMoveOne, ok := fuelToMoveOneCache[diff]
			if !ok {
				n := 0
				for i := 0; i <= diff; i++ {
					n += i
				}
				fuelToMoveOne = n
				fuelToMoveOneCache[diff] = n
			}

			totalFuelToPos += fuelToMoveOne * count
		}

		fuelToPosition[pos] = totalFuelToPos
	}

	least := -1
	for _, fuel := range fuelToPosition {
		if least == -1 {
			least = fuel
		}

		if fuel < least {
			least = fuel
		}
	}

	return least
}

func TestDaySeven(t *testing.T) {
	type testConfig struct {
		input     []int
		expected  int
		logResult bool
	}

	input, err := readFileAsCommaSeparatedInts("./input/07.txt")
	if err != nil {
		t.Fatalf("failed to read input")
	}

	runTests := func(t *testing.T, tests map[string]testConfig, fn func([]int) int) {
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
				input: []int{
					16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
				},
				expected: 37,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, theTreacheryOfWhalesPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test 1": {
				input: []int{
					16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
				},
				expected: 168,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, theTreacheryOfWhalesPartTwo)
	})
}
