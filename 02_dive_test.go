package aoc

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func divePartOne(commands []string) int {
	horizontal := 0
	depth := 0

	for _, command := range commands {
		split := strings.Split(command, " ")

		n, err := strconv.Atoi(split[1])
		if err != nil {
			panic(fmt.Sprintf("strconv.Atoi: %v", split[1]))
		}

		switch split[0] {
		case "forward":
			horizontal += n
		case "down":
			depth += n
		case "up":
			depth -= n
		}
	}

	return horizontal * depth
}

func divePartTwo(commands []string) int {
	horizontal := 0
	depth := 0
	aim := 0

	for _, command := range commands {
		split := strings.Split(command, " ")

		n, err := strconv.Atoi(split[1])
		if err != nil {
			panic(fmt.Sprintf("strconv.Atoi: %v", split[1]))
		}

		switch split[0] {
		case "forward":
			horizontal += n
			depth += (aim * n)
		case "down":
			aim += n
		case "up":
			aim -= n
		}
	}

	return horizontal * depth
}

func TestDayTwo(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/02.txt")
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
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, divePartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, divePartTwo)
	})
}
