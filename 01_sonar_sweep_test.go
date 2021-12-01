package aoc

import (
	"testing"
)

func sonarSweepPartOne(measurements []int) int {
	total := 0
	prev := -1
	for _, measurement := range measurements {
		if prev == -1 {
			prev = measurement
			continue
		}

		if measurement > prev {
			total += 1
		}
		prev = measurement
	}

	return total
}

func sonarSweepPartTwo(measurements []int) int {
	total := 0

	prev := measurements[0] + measurements[1] + measurements[2]
	for i := 3; i < len(measurements); i++ {
		curr := measurements[i] + measurements[i-1] + measurements[i-2]

		if curr > prev {
			total += 1
		}

		prev = curr
	}

	return total
}

func TestDayOne(t *testing.T) {
	type testConfig struct {
		input     []int
		expected  int
		logResult bool
	}

	input, err := readFileAsIntSlice("./input/01.txt")
	if err != nil {
		t.Fatalf("failed to read input")
	}

	runTests := func(t *testing.T, tests map[string]testConfig, fn func([]int) int) {
		for name, cfg := range tests {
			cfg := cfg
			t.Run(name, func(t *testing.T) {
				output := fn(cfg.input)
				if cfg.logResult {
					t.Log(output)
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
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, sonarSweepPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test 1": {
				input: []int{
					199,
					200,
					208,
					210,
					200,
					207,
					240,
					269,
					260,
					263,
				},
				expected: 5,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, sonarSweepPartTwo)
	})
}
