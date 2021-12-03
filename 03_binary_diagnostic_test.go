package aoc

import (
	"fmt"
	"testing"
	"time"
)

func binaryDiagnosticPartOne(numbers []string) int {
	significantBits := make([]int, len(numbers[0]))

	for _, number := range numbers {
		for i, bit := range number {
			switch bit {
			case '0':
				significantBits[i] -= 1
			case '1':
				significantBits[i] += 1
			}
		}
	}

	gamma, epsilon := 0, 0

	for _, sigBit := range significantBits {
		gamma = gamma << 1
		epsilon = epsilon << 1

		if sigBit > 0 {
			gamma += 1
		} else {
			epsilon += 1
		}
	}

	return gamma * epsilon
}

func binaryDiagnosticPartTwo(numbers []string) int {
	return 0
}

func TestDayThree(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/03.txt")
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
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
				expected: 198,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, binaryDiagnosticPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test 1": {
				input: []string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
				expected: 198,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, binaryDiagnosticPartTwo)
	})
}
