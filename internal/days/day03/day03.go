package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Strictly match mul(X,Y) where X and Y are 1-3 digits with no spaces or extra chars.
var mulRE = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func getMultiplyTotal(line string) int {
	total := 0
	matches := mulRE.FindAllStringSubmatch(line, -1)
	for _, m := range matches {
		// m[1] and m[2] are guaranteed to be 1-3 digits each by the regex
		a, errA := strconv.Atoi(m[1])
		if errA != nil {
			continue
		}
		b, errB := strconv.Atoi(m[2])
		if errB != nil {
			continue
		}
		// Bounds implicitly <= 999 due to {1,3} digits; no negatives possible
		total += a * b
	}
	return total
}

func SolvePart1(input string) (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("failed to close input file: %v\n", cerr)
		}
	}()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	if len(lines) == 0 {
		return 0, fmt.Errorf("no input data found")
	}
	multiply_total := 0
	for _, line := range lines {
		multiply_total += getMultiplyTotal(line)
	}
	return multiply_total, nil
}

const (
	doStr   = "do()"
	dontStr = "don't()"
)

func getMultiplyTotalForDo(line string) int {
	// The do() instruction enables future mul instructions.
	// The don't() instruction disables future mul instructions.
	// Only the most recent do() or don't() instruction applies. At the beginning of the program, mul instructions are enabled.
	total := 0
	parts := strings.Split(line, doStr)
	for _, part := range parts {
		subParts := strings.Split(part, dontStr)
		if len(subParts) > 0 {
			total += getMultiplyTotal(subParts[0])
		} else {
			total += getMultiplyTotal(part)
		}
	}
	return total
}

func SolvePart2(input string) (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("failed to close input file: %v\n", cerr)
		}
	}()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	if len(lines) == 0 {
		return 0, fmt.Errorf("no input data found")
	}
	// Combine lines into a single string
	combined := strings.Join(lines, " ")
	multiply_total := getMultiplyTotalForDo(combined)
	return multiply_total, nil
}
