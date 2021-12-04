package aoc

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

type bingoGame struct {
	boards []bingoBoard
}

func (g *bingoGame) addBoard(board bingoBoard) {
	g.boards = append(g.boards, board)
}

func (g *bingoGame) callNumber(num string) (bool, bingoBoard) {
	for _, board := range g.boards {
		winner := board.markNumber(num)
		if winner {
			return true, board
			break
		}
	}

	return false, bingoBoard{}
}

type bingoBoard struct {
	rows, columns, added int
	numbers              [][]bingoNumber
	rev                  map[string][]int
}

func (b *bingoBoard) init() {
	b.rows = 5
	b.columns = 5
	b.numbers = make([][]bingoNumber, 5, 5)
	for i := 0; i < 5; i++ {
		b.numbers[i] = make([]bingoNumber, 5, 5)
	}
	b.rev = make(map[string][]int)
}

func (b *bingoBoard) addNumber(number bingoNumber) {
	x, y := 0, 0

	x = b.added % b.columns

	for i := 1; i < b.rows+1; i++ {
		if b.added >= b.rows*i {
			y++
		}
	}

	b.numbers[y][x] = number
	b.rev[number.value] = []int{x, y}

	b.added += 1
}

func (b *bingoBoard) markNumber(number string) bool {
	loc, ok := b.rev[number]
	if !ok {
		return false
	}

	x := loc[0]
	y := loc[1]

	b.numbers[y][x].marked = true

	allMarked := true
	for i := 0; i < b.rows; i++ {
		if !b.numbers[i][x].marked {
			allMarked = false
		}
	}
	if allMarked {
		return true
	}

	allMarked = true
	for i := 0; i < b.columns; i++ {
		if !b.numbers[y][i].marked {
			allMarked = false
		}
	}

	return allMarked
}

func (b *bingoBoard) sumUnmarked() int {
	total := 0

	for _, row := range b.numbers {
		for _, num := range row {
			if !num.marked {
				n, err := strconv.Atoi(num.value)
				if err != nil {
					panic(fmt.Sprintf("error converting unmarked number %s to int", num.value))
				}
				total += n
			}
		}
	}

	return total
}

type bingoNumber struct {
	value  string
	marked bool
}

func giantSquidPartOne(bingoNums []string) int {
	allNumbers := bingoNums[0]
	callNumbers := strings.Split(allNumbers, ",")

	game := bingoGame{
		boards: []bingoBoard{},
	}

	board := bingoBoard{}
	board.init()

	for i := 2; i < len(bingoNums); i++ {
		if bingoNums[i] == "" {
			game.addBoard(board)
			board = bingoBoard{}
			board.init()
			continue
		}

		boardNumbers := strings.Split(bingoNums[i], " ")

		for _, n := range boardNumbers {
			if n == "" {
				continue
			}

			bingoNum := bingoNumber{
				value:  n,
				marked: false,
			}
			board.addNumber(bingoNum)
		}
	}
	game.addBoard(board)

	winningCall := ""
	var winningBoard bingoBoard
	for _, call := range callNumbers {
		winner, board := game.callNumber(call)
		if winner {
			winningCall = call
			winningBoard = board
			break
		}
	}

	unmarkedTotal := winningBoard.sumUnmarked()
	winningCallNum, err := strconv.Atoi(winningCall)
	if err != nil {
		panic("error converting winningCallNum to int")
	}

	return winningCallNum * unmarkedTotal
}

func TestDayFour(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/04.txt")
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
					"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
					"",
					"22 13 17 11  0",
					" 8  2 23  4 24",
					"21  9 14 16  7",
					" 6 10  3 18  5",
					" 1 12 20 15 19",
					"",
					" 3 15  0  2 22",
					" 9 18 13 17  5",
					"19  8  7 25 23",
					"20 11 10 24  4",
					"14 21 16 12  6",
					"",
					"14 21 17 24  4",
					"10 16 15  9 19",
					"18  8 23 26 20",
					"22 11 13  6  5",
					" 2  0 12  3  7",
				},
				expected: 4512,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, giantSquidPartOne)
	})
}
