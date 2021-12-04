package aoc

import (
	"fmt"
	"strconv"
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
	binaryNumLen := len(numbers[0])

	var firstBit0, firstBit1 []string
	var oxygenNums, co2Nums []string

	for _, number := range numbers {
		switch number[0] {
		case '0':
			firstBit0 = append(firstBit0, number)
		case '1':
			firstBit1 = append(firstBit1, number)
		}
	}

	if len(firstBit0) > len(firstBit1) {
		oxygenNums = firstBit0
		co2Nums = firstBit1
	} else {
		oxygenNums = firstBit1
		co2Nums = firstBit0
	}

	ratings := make(chan string, 2)

	getRating := func(nums []string, greedy bool) {
		for i := 1; i <= binaryNumLen; i++ {
			if len(nums) == 1 {
				ratings <- nums[0]
				break
			}

			var tmp0, tmp1 []string

			for _, number := range nums {
				switch number[i] {
				case '0':
					tmp0 = append(tmp0, number)
				case '1':
					tmp1 = append(tmp1, number)
				}

				if len(tmp0) == len(tmp1) {
					if tmp0[0][i] == '1' {
						if greedy {
							nums = tmp0
						} else {
							nums = tmp1
						}
					} else {
						if greedy {
							nums = tmp1
						} else {
							nums = tmp0
						}
					}
				} else if len(tmp0) > len(tmp1) {
					if greedy {
						nums = tmp0
					} else {
						nums = tmp1
					}
				} else {
					if greedy {
						nums = tmp1
					} else {
						nums = tmp0
					}
				}
			}
		}
	}

	go getRating(oxygenNums, true)
	go getRating(co2Nums, false)

	num1 := <-ratings
	num2 := <-ratings

	num1Int, err := strconv.ParseInt(num1, 2, 64)
	if err != nil {
		panic(fmt.Sprintf("cant convert binary num %s to int", num1))
	}
	num2Int, err := strconv.ParseInt(num2, 2, 64)
	if err != nil {
		panic(fmt.Sprintf("cant convert binary num %s to int", num2))
	}

	return int(num1Int * num2Int)
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
				expected: 230,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, binaryDiagnosticPartTwo)
	})
}
