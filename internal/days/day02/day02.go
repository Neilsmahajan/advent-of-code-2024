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
		status := "unsafe"
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
	// TODO: Implement solution for Part 2
	return 0, fmt.Errorf("day 02 part 2 not implemented yet")
}
