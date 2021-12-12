package aoc

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

type caveStack struct {
	values []string
}

func (s *caveStack) push(node string) {
	s.values = append(s.values, node)
}

func (s *caveStack) peek() string {
	lastIndex := len(s.values) - 1
	return s.values[lastIndex]
}

func (s *caveStack) contains(target string) bool {
	for _, v := range s.values {
		if v == target {
			return true
		}
	}
	return false
}

func (s *caveStack) count(target string) int {
	total := 0
	for _, v := range s.values {
		if v == target {
			total++
		}
	}
	return total
}

func (s *caveStack) getValues() []string {
	return s.values
}

func (s *caveStack) clone(source *caveStack) {
	copyValues := source.getValues()
	for _, v := range copyValues {
		s.values = append(s.values, v)
	}
}

func passagePathingPartOne(joins []string) int {
	caves := map[string][]string{}

	for _, join := range joins {
		split := strings.Split(join, "-")
		caves[split[0]] = append(caves[split[0]], split[1])
		caves[split[1]] = append(caves[split[1]], split[0])
	}

	stack := &caveStack{}
	stack.push("start")
	total := 0
	paths := []*caveStack{stack}
	for len(paths) > 0 {
		newPaths := []*caveStack{}
		finishedPaths := []int{}

		for i, path := range paths {
			node := path.peek()

			if node == "end" {
				total++
				finishedPaths = append(finishedPaths, i)
				continue
			}

			deadEnd := false
			if strings.ToLower(node) == node {
				if path.count(node) > 1 {
					deadEnd = true
				}
			}
			if node == "start" && path.count(node) > 1 {
				deadEnd = true
			}

			if deadEnd {
				finishedPaths = append(finishedPaths, i)
				continue
			}

			options := caves[node]

			if len(options) > 1 {
				for j := 1; j < len(options); j++ {
					opt := options[j]
					newPath := &caveStack{}
					newPath.clone(path)
					newPath.push(opt)
					newPaths = append(newPaths, newPath)
				}
			}

			// always add the first option to the current path
			if len(options) > 0 {
				firstOpt := options[0]
				path.push(firstOpt)
			}
		}

		// remove ended paths
		paths = removeItems(paths, finishedPaths)

		// add new paths
		paths = append(paths, newPaths...)
	}

	return total
}

// remove given indexes from slice
func removeItems(s []*caveStack, toRemove []int) []*caveStack {
	for _, remove := range toRemove {
		s[remove] = nil
	}

	updated := []*caveStack{}
	for _, v := range s {
		if v != nil {
			updated = append(updated, v)
		}
	}

	return updated
}

func TestDayTwelve(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/12.txt")
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
					"start-A",
					"start-b",
					"A-c",
					"A-b",
					"b-d",
					"A-end",
					"b-end",
				},
				expected: 10,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, passagePathingPartOne)
	})
}
