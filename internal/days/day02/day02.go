package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SolvePart1(input string) (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return 0, fmt.Errorf("failed to open input file: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("failed to close input file: %v\n", cerr)
		}
	}()
	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		var report []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return 0, fmt.Errorf("failed to parse number %s: %w", part, err)
			}
			report = append(report, num)
		}
		if len(report) > 0 {
			reports = append(reports, report)
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to read input file: %w", err)
	}
	safeCount := 0
	for _, report := range reports {
		var status string
		if report[1] > report[0] && report[1] <= report[0]+3 {
			status = "increasing"
		} else if report[1] < report[0] && report[1] >= report[0]-3 {
			status = "decreasing"
		} else {
			continue
		}
		for level := 2; level < len(report); level++ {
			switch status {
			case "increasing":
				if report[level] <= report[level-1] || report[level] > report[level-1]+3 {
					status = "unsafe"
				}
			case "decreasing":
				if report[level] >= report[level-1] || report[level] < report[level-1]-3 {
					status = "unsafe"
				}
			}
		}
		if status != "unsafe" {
			safeCount++
		}
	}
	return safeCount, nil
}

func SolvePart2(input string) (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return 0, fmt.Errorf("failed to open input file: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("failed to close input file: %v\n", cerr)
		}
	}()
	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		var report []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return 0, fmt.Errorf("failed to parse number %s: %w", part, err)
			}
			report = append(report, num)
		}
		if len(report) > 0 {
			reports = append(reports, report)
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to read input file: %w", err)
	}

	safeCount := 0
	for _, report := range reports {
		if isSafeWithDampener(report) {
			safeCount++
		}
	}

	return safeCount, nil
}

func isSafeWithDampener(report []int) bool {
	// First check if it's already safe without removing anything
	if isSafeReport(report) {
		return true
	}

	// Try removing each level one by one
	for i := 0; i < len(report); i++ {
		// Create a new slice without the element at index i
		modified := make([]int, 0, len(report)-1)
		modified = append(modified, report[:i]...)
		modified = append(modified, report[i+1:]...)

		if isSafeReport(modified) {
			return true
		}
	}

	return false
}

func isSafeReport(report []int) bool {
	if len(report) < 2 {
		return true
	}

	// Determine if sequence should be increasing or decreasing
	increasing := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Check if direction is consistent
		if increasing && diff <= 0 {
			return false
		}
		if !increasing && diff >= 0 {
			return false
		}

		// Check if difference is within valid range (1-3)
		absDiff := diff
		if absDiff < 0 {
			absDiff = -absDiff
		}
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}

	return true
}
