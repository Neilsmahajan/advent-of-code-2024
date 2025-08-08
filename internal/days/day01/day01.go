package day01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/neilsmahajan/advent-of-code-2024/internal/utils"
)

func SolvePart1(input string) (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()
	var leftList, rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // Splits by whitespace
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid line format: %s", line)
		}

		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, fmt.Errorf("invalid left number: %s", parts[0])
		}

		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, fmt.Errorf("invalid right number: %s", parts[1])
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading from file: %v", err)
	}

	// Ensure the lists can be paired up.
	if len(leftList) != len(rightList) {
		return 0, fmt.Errorf("the lists have different lengths and cannot be paired")
	}

	// 1. Sort both lists to align numbers for pairing.
	sort.Ints(leftList)
	sort.Ints(rightList)

	var totalDistance int64 // Use int64 for the sum to avoid potential overflow.

	// 2. Calculate the sum of absolute differences for each pair.
	for i := 0; i < len(leftList); i++ {
		distance := utils.AbsInt(leftList[i] - rightList[i])
		totalDistance += int64(distance)
	}
	return int(totalDistance), nil
}

func SolvePart2(input string) (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()
	var leftList, rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // Splits by whitespace
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid line format: %s", line)
		}

		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, fmt.Errorf("invalid left number: %s", parts[0])
		}

		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, fmt.Errorf("invalid right number: %s", parts[1])
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading from file: %v", err)
	}

	// Count frequency of each number in the right list
	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	// Calculate similarity score
	var similarityScore int64
	for _, leftNum := range leftList {
		count := rightCount[leftNum]
		similarityScore += int64(leftNum * count)
	}

	return int(similarityScore), nil
}
