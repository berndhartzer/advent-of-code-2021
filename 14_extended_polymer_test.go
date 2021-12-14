package aoc

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func extendedPolymerizationPartOne(input []string) int {
	return extendedPolymerization(input, 10)
}

func extendedPolymerizationPartTwo(input []string) int {
	return extendedPolymerization(input, 40)
}

func extendedPolymerization(input []string, steps int) int {
	pairs := map[[2]byte]int{}
	chars := map[byte]int{}
	start := input[0]
	for j := 0; j < len(start)-1; j++ {
		pairs[[2]byte{start[j], start[j+1]}]++
		chars[start[j]]++
	}
	chars[start[len(start)-1]]++

	k := 2
	rules := map[[2]byte]byte{}
	for ; k < len(input); k++ {
		split := strings.Split(input[k], " -> ")
		rules[[2]byte{byte(split[0][0]), byte(split[0][1])}] = byte(split[1][0])
	}

	for i := 0; i < steps; i++ {
		delta := map[[2]byte]int{}

		for pair, count := range pairs {
			ruleRes, ok := rules[pair]
			if ok {
				delta[pair] -= 1 * count
				left, right := pair[0], pair[1]

				delta[[2]byte{left, ruleRes}] += 1 * count
				delta[[2]byte{ruleRes, right}] += 1 * count
				chars[ruleRes] += 1 * count
			}
		}

		for pair, count := range delta {
			pairs[pair] += count
		}
	}

	highest, lowest := 0, -1
	for _, v := range chars {
		if v > highest {
			highest = v
		}
		if lowest == -1 || v < lowest {
			lowest = v
		}
	}

	return highest - lowest
}

func TestDayFourteen(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/14.txt")
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
					"NNCB",
					"",
					"CH -> B",
					"HH -> N",
					"CB -> H",
					"NH -> C",
					"HB -> C",
					"HC -> B",
					"HN -> C",
					"NN -> C",
					"BH -> H",
					"NC -> B",
					"NB -> B",
					"BN -> B",
					"BB -> N",
					"BC -> B",
					"CC -> N",
					"CN -> C",
				},
				expected: 1588,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, extendedPolymerizationPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, extendedPolymerizationPartTwo)
	})
}
