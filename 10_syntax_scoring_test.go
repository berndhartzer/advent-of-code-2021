package aoc

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

const (
	openParens  = rune('(')
	closeParens = rune(')')
	openBrace   = rune('{')
	closeBrace  = rune('}')
	openSquare  = rune('[')
	closeSquare = rune(']')
	openAngle   = rune('<')
	closeAngle  = rune('>')
)

var closer = map[rune]rune{
	closeParens: openParens,
	closeBrace:  openBrace,
	closeSquare: openSquare,
	closeAngle:  openAngle,
}
var opener = map[rune]rune{
	openParens: closeParens,
	openBrace:  closeBrace,
	openSquare: closeSquare,
	openAngle:  closeAngle,
}

type syntaxStack struct {
	values []rune
}

func (s *syntaxStack) push(v rune) {
	s.values = append(s.values, v)
}

func (s *syntaxStack) pop() rune {
	lastIndex := len(s.values) - 1
	v := s.values[lastIndex]
	s.values = s.values[:lastIndex]
	return v
}

func (s *syntaxStack) length() int {
	return len(s.values)
}

func syntaxScoringPartOne(lines []string) int {
	scorer := map[rune]int{
		closeParens: 3,
		closeBrace:  1197,
		closeSquare: 57,
		closeAngle:  25137,
	}

	score := 0

	for _, line := range lines {
		stack := syntaxStack{}

		for _, c := range line {
			switch c {
			case openParens:
				fallthrough
			case openBrace:
				fallthrough
			case openSquare:
				fallthrough
			case openAngle:
				stack.push(c)
			case closeParens:
				fallthrough
			case closeBrace:
				fallthrough
			case closeSquare:
				fallthrough
			case closeAngle:
				head := stack.pop()
				expected := closer[c]
				if head != expected {
					score += scorer[c]
					break
				}
			}
		}
	}

	return score
}

func syntaxScoringPartTwo(lines []string) int {
	scorer := map[rune]int{
		openParens: 1,
		openBrace:  3,
		openSquare: 2,
		openAngle:  4,
	}

	lineScores := []int{}

lineLooper:
	for _, line := range lines {
		stack := syntaxStack{}

		for _, c := range line {
			switch c {
			case openParens:
				fallthrough
			case openBrace:
				fallthrough
			case openSquare:
				fallthrough
			case openAngle:
				stack.push(c)
			case closeParens:
				fallthrough
			case closeBrace:
				fallthrough
			case closeSquare:
				fallthrough
			case closeAngle:
				head := stack.pop()
				expected := closer[c]
				if head != expected {
					continue lineLooper
				}
			}
		}

		lineScore := 0
		for stack.length() > 0 {
			lineScore *= 5
			lastOpen := stack.pop()
			lineScore += scorer[lastOpen]
		}

		lineScores = append(lineScores, lineScore)
	}

	sort.Ints(lineScores)
	return lineScores[(len(lineScores) / 2)]
}

func TestDayTen(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/10.txt")
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
					"[({(<(())[]>[[{[]{<()<>>",
					"[(()[<>])]({[<{<<[]>>(",
					"{([(<{}[<>[]}>{[]{[(<()>",
					"(((({<>}<{<{<>}{[]{[]{}",
					"[[<[([]))<([[{}[[()]]]",
					"[{[{({}]{}}([{[{{{}}([]",
					"{<[[]]>}<{[{[{[]{()[[[]",
					"[<(<(<(<{}))><([]([]()",
					"<{([([[(<>()){}]>(<<{{",
					"<{([{{}}[<[[[<>{}]]]>[]]",
				},
				expected: 26397,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, syntaxScoringPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test 1": {
				input: []string{
					"[({(<(())[]>[[{[]{<()<>>",
					"[(()[<>])]({[<{<<[]>>(",
					"{([(<{}[<>[]}>{[]{[(<()>",
					"(((({<>}<{<{<>}{[]{[]{}",
					"[[<[([]))<([[{}[[()]]]",
					"[{[{({}]{}}([{[{{{}}([]",
					"{<[[]]>}<{[{[{[]{()[[[]",
					"[<(<(<(<{}))><([]([]()",
					"<{([([[(<>()){}]>(<<{{",
					"<{([{{}}[<[[[<>{}]]]>[]]",
				},
				expected: 288957,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, syntaxScoringPartTwo)
	})
}
