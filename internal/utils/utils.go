package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads all lines from a file and returns them as a slice of strings
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ReadInts reads all lines from a file and converts them to integers
func ReadInts(filename string) ([]int, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	var ints []int
	for _, line := range lines {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, fmt.Errorf("invalid integer: %s", line)
		}
		ints = append(ints, num)
	}
	return ints, nil
}

// SplitInts splits a string by delimiter and converts to integers
func SplitInts(s string, delimiter string) ([]int, error) {
	parts := strings.Split(s, delimiter)
	var ints []int
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid integer: %s", part)
		}
		ints = append(ints, num)
	}
	return ints, nil
}

// AbsInt returns the absolute value of an integer
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// MinInt returns the minimum of two integers
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the maximum of two integers
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// SumInts returns the sum of a slice of integers
func SumInts(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// ParseGrid parses a 2D character grid from lines
func ParseGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

// GridBounds returns the dimensions of a 2D grid
func GridBounds(grid [][]rune) (rows, cols int) {
	if len(grid) == 0 {
		return 0, 0
	}
	return len(grid), len(grid[0])
}

// InBounds checks if coordinates are within grid bounds
func InBounds(grid [][]rune, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}
