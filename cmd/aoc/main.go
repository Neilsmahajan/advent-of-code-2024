package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/neilsmahajan/advent-of-code-2024/internal/days/day01"
	// "github.com/neilsmahajan/advent-of-code-2024/internal/days/day02"
)

type Solution struct {
	Part1 func(string) (int, error)
	Part2 func(string) (int, error)
}

var solutions = map[int]Solution{
	1: {day01.SolvePart1, day01.SolvePart2},
	// 2: {day02.SolvePart1, day02.SolvePart2},
}

func main() {
	var day int
	var part int
	var all bool

	flag.IntVar(&day, "day", 1, "Day to solve (1-25)")
	flag.IntVar(&part, "part", 1, "Part to solve (1 or 2)")
	flag.BoolVar(&all, "all", false, "Run all implemented solutions")
	flag.Parse()

	if all {
		runAllSolutions()
		return
	}

	if day < 1 || day > 25 {
		fmt.Printf("Error: Day must be between 1 and 25, got %d\n", day)
		os.Exit(1)
	}

	if part < 1 || part > 2 {
		fmt.Printf("Error: Part must be 1 or 2, got %d\n", part)
		os.Exit(1)
	}

	inputPath := filepath.Join("internal", "days", fmt.Sprintf("day%02d", day), "input.txt")

	solution, exists := solutions[day]
	if !exists {
		fmt.Printf("Error: Day %d not implemented yet\n", day)
		os.Exit(1)
	}

	var result int
	var err error

	if part == 1 {
		result, err = solution.Part1(inputPath)
	} else {
		result, err = solution.Part2(inputPath)
	}

	if err != nil {
		fmt.Printf("Error solving Day %d Part %d: %v\n", day, part, err)
		os.Exit(1)
	}

	fmt.Printf("Day %d Part %d: %d\n", day, part, result)
}

func runAllSolutions() {
	fmt.Println("Running all implemented solutions:")
	fmt.Println("===================================")

	for day := 1; day <= 25; day++ {
		solution, exists := solutions[day]
		if !exists {
			continue
		}

		inputPath := filepath.Join("internal", "days", fmt.Sprintf("day%02d", day), "input.txt")

		// Check if input file exists
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			fmt.Printf("Day %02d: Input file not found\n", day)
			continue
		}

		// Try Part 1
		if result, err := solution.Part1(inputPath); err == nil {
			fmt.Printf("Day %02d Part 1: %d\n", day, result)
		} else {
			fmt.Printf("Day %02d Part 1: Error - %v\n", day, err)
		}

		// Try Part 2
		if result, err := solution.Part2(inputPath); err == nil {
			fmt.Printf("Day %02d Part 2: %d\n", day, result)
		} else {
			fmt.Printf("Day %02d Part 2: Error - %v\n", day, err)
		}

		fmt.Println()
	}
}
