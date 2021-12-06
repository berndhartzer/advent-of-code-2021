package aoc

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readFileAsIntSlice(name string) ([]int, error) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var contents []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		contents = append(contents, num)
	}

	return contents, nil
}

func readFileAsStringSlice(name string) ([]string, error) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var contents []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents, nil
}

func readFileAsString(name string) (string, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}

	s := string(b)
	s = strings.TrimSuffix(s, "\n")

	return s, nil
}

func readFileAsCommaSeparatedInts(name string) ([]int, error) {
	stringVal, err := readFileAsString(name)
	if err != nil {
		return nil, nil
	}
	split := strings.Split(stringVal, ",")

	intVals := []int{}
	for _, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, nil
		}

		intVals = append(intVals, n)
	}

	return intVals, nil
}
