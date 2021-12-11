package aoc

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func sevenSegmentSearchPartOne(signals []string) int {
	unique := 0
	for _, signal := range signals {
		split := strings.Split(signal, "|")
		outputValues := split[1]
		outputSplit := strings.Split(outputValues, " ")
		for _, value := range outputSplit {
			switch len(value) {
			case 2: // 1
				unique++
			case 4: // 4
				unique++
			case 3: // 7
				unique++
			case 7: // 8
				unique++
			}
		}
	}
	return unique
}

type digit struct {
	topMiddle    string
	topLeft      string
	topRight     string
	middle       string
	bottomLeft   string
	bottomRight  string
	bottomMiddle string
}

func (d digit) getChars() string {
	res := ""
	res += d.topMiddle
	res += d.topLeft
	res += d.topRight
	res += d.middle
	res += d.bottomLeft
	res += d.bottomRight
	res += d.bottomMiddle
	return res
}

func (d digit) asInt(s string) int {
	switch len(s) {
	case 2:
		return 1
	case 4:
		return 4
	case 3:
		return 7
	case 7:
		return 8
	case 5:
		if strings.Contains(s, d.topLeft) {
			return 5
		}
		if strings.Contains(s, d.bottomLeft) {
			return 2
		}
		return 3
	case 6:
		if !strings.Contains(s, d.middle) {
			return 0
		}
		if strings.Contains(s, d.bottomLeft) {
			return 6
		}
		return 9
	}

	panic("we shouldnt be here")
}

// pretty ugly solution, but it works, and ive got cricket to watch
func sevenSegmentSearchPartTwo(signals []string) int {
	total := 0

	for _, signal := range signals {
		split := strings.Split(signal, " ")
		inputs := split[:10]
		outputs := split[len(split)-4:]

		digits := map[int]string{}

		d := digit{}
		lenFives := []string{}
		lenSixes := []string{}

		for _, input := range inputs {
			switch len(input) {
			case 2:
				digits[1] = input
			case 4:
				digits[4] = input
			case 3:
				digits[7] = input
			case 7:
				digits[8] = input
			case 5:
				lenFives = append(lenFives, input)
			case 6:
				lenSixes = append(lenSixes, input)

			}
		}

		d.topMiddle = getUniqueChars(digits[1], digits[7])

		middleAndTopLeft := getUniqueChars(digits[1], digits[4])
		for _, s := range lenSixes {
			matching := getMatchingChars(s, middleAndTopLeft)
			if len(matching) == 1 {
				d.topLeft = matching[0]
				digits[0] = s
				break
			}
		}
		d.middle = getUniqueChars(d.topLeft, middleAndTopLeft)

		bottomMiddleAndRight := ""
		for _, s := range lenFives {
			currentKnown := d.getChars()
			matching := getMatchingChars(s, currentKnown)
			if len(matching) == 3 {
				bottomMiddleAndRight = getUniqueChars(s, currentKnown)
				digits[5] = s
			}
		}

		d.bottomRight = getMatchingChars(bottomMiddleAndRight, digits[1])[0]
		d.topRight = getUniqueChars(digits[1], d.bottomRight)
		d.bottomMiddle = getUniqueChars(bottomMiddleAndRight, d.bottomRight)
		d.bottomLeft = getUniqueChars(d.getChars(), "abcdefg")

		x := 1000
		for _, output := range outputs {
			asInt := d.asInt(output)
			total += x * asInt

			x /= 10
		}
	}

	return total
}

func getMatchingChars(a, b string) []string {
	charMap := getCharMap(a, b)
	result := []string{}
	for char, c := range charMap {
		if c > 1 {
			result = append(result, string(char))
		}
	}

	return result
}

func getUniqueChars(a, b string) string {
	charMap := getCharMap(a, b)
	result := ""
	for char, c := range charMap {
		if c == 1 {
			result += string(char)
		}
	}

	return result
}

func getCharMap(a, b string) map[rune]int {
	chars := map[rune]int{}

	for _, char := range a {
		_, ok := chars[char]
		if ok {
			chars[char]++
		} else {
			chars[char] = 1
		}
	}
	for _, char := range b {
		_, ok := chars[char]
		if ok {
			chars[char]++
		} else {
			chars[char] = 1
		}
	}

	return chars
}

func TestDayEight(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/08.txt")
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
					"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
					"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
					"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
					"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
					"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
					"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
					"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
					"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
					"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
					"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
				},
				expected: 26,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, sevenSegmentSearchPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test 1": {
				input: []string{
					"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf",
				},
				expected: 5353,
			},
			"test 2": {
				input: []string{
					"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
					"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
					"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
					"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
					"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
					"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
					"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
					"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
					"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
					"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
				},
				expected: 61229,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, sevenSegmentSearchPartTwo)
	})
}
