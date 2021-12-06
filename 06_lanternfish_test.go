package aoc

import (
	"fmt"
	"testing"
	"time"
)

func lanternfishPartOne(fish []int, days int) int {
	for day := 0; day < days; day++ {
		newFish := []int{}

		for i := range fish {
			fish[i]--
			if fish[i] < 0 {
				fish[i] = 6
				newFish = append(newFish, 8)
			}
		}

		fish = append(fish, newFish...)
	}

	return len(fish)
}

func lanternfishPartTwo(fish []int, days int) int {
	fishOnDays := make([]int, 9)

	for _, f := range fish {
		fishOnDays[f]++
	}

	for day := 0; day < days; day++ {
		numNewFish := fishOnDays[0]
		for i := 0; i < len(fishOnDays)-1; i++ {
			fishOnDays[i] = fishOnDays[i+1]
		}

		fishOnDays[8] = numNewFish
		fishOnDays[6] += numNewFish
	}

	totalFish := 0
	for _, fish := range fishOnDays {
		totalFish += fish
	}

	return totalFish
}

func TestDaySix(t *testing.T) {
	type testConfig struct {
		input     []int
		days      int
		expected  int
		logResult bool
	}

	input, err := readFileAsCommaSeparatedInts("./input/06.txt")
	if err != nil {
		t.Fatalf("failed to read input")
	}

	runTests := func(t *testing.T, tests map[string]testConfig, fn func([]int, int) int) {
		for name, cfg := range tests {
			cfg := cfg
			t.Run(name, func(t *testing.T) {
				start := time.Now()
				output := fn(cfg.input, cfg.days)
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
					3, 4, 3, 1, 2,
				},
				days:     18,
				expected: 26,
			},
			"test 2": {
				input: []int{
					3, 4, 3, 1, 2,
				},
				days:     80,
				expected: 5934,
			},
			"solution": {
				input:     input,
				days:      80,
				logResult: true,
			},
		}

		runTests(t, tests, lanternfishPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test 1": {
				input: []int{
					3, 4, 3, 1, 2,
				},
				days:     256,
				expected: 26984457539,
			},
			"solution": {
				input:     input,
				days:      256,
				logResult: true,
			},
		}

		runTests(t, tests, lanternfishPartTwo)
	})
}
