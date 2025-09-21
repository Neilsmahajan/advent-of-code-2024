package day05

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// SolvePart1 computes the sum of the middle numbers of all valid update lists
// according to the ordering rules defined at the top of the input file.
func SolvePart1(input string) (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("failed to close input file: %v\n", err)
		}
	}()

	mustComeBefore, updates, err := parseRulesAndLists(file)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, nums := range updates {
		valid, _, _ := isValidUpdate(nums, mustComeBefore)
		if valid {
			total += middleValue(nums)
		}
	}
	return total, nil
}

// parseRulesAndLists reads the input from r and returns the ordering rules and the update lists.
// It expects a block of pipe-delimited rules (A|B meaning A must come before B) followed by a blank line,
// and then comma-delimited update lines.
func parseRulesAndLists(r io.Reader) (map[int]map[int]bool, [][]int, error) {
	mustComeBefore := make(map[int]map[int]bool)
	scanner := bufio.NewScanner(r)

	// Parse rules section.
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" { // blank line separates rules from updates
			break
		}
		numbers := strings.Split(line, "|")
		if len(numbers) != 2 {
			return nil, nil, fmt.Errorf("invalid input line: %s", line)
		}

		before, err := strconv.Atoi(numbers[0])
		if err != nil {
			return nil, nil, err
		}
		after, err := strconv.Atoi(numbers[1])
		if err != nil {
			return nil, nil, err
		}
		beforeMap := mustComeBefore[after]
		if beforeMap == nil {
			beforeMap = make(map[int]bool)
		}
		beforeMap[before] = true
		mustComeBefore[after] = beforeMap
	}

	// Parse update lists.
	updates := make([][]int, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" { // ignore stray blank lines below
			continue
		}
		numStrings := strings.Split(line, ",")
		nums := make([]int, 0, len(numStrings))
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return nil, nil, err
			}
			nums = append(nums, num)
		}
		updates = append(updates, nums)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return mustComeBefore, updates, nil
}

// isValidUpdate returns true if every number that must appear before another does so
// within the provided nums slice.
func isValidUpdate(nums []int, mustComeBefore map[int]map[int]bool) (bool, int, int) {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			after := nums[j]
			beforeSet := mustComeBefore[after]
			if beforeSet[num] == false { // either absent or explicitly false
				return false, i, j
			}
		}
	}
	return true, -1, -1
}

// middleValue returns the middle element of nums. Assumes nums has odd length per puzzle input.
func middleValue(nums []int) int {
	return nums[len(nums)/2]
}

func SolvePart2(input string) (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("failed to close input file: %v\n", err)
		}
	}()

	mustComeBefore, updates, err := parseRulesAndLists(file)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, nums := range updates {
		valid, _, _ := isValidUpdate(nums, mustComeBefore)
		if valid {
			continue
		}
		reorderedNums := reorderToValid(nums, mustComeBefore)
		total += middleValue(reorderedNums)
	}
	return total, nil
}

func reorderToValid(nums []int, mustComeBefore map[int]map[int]bool) []int {
	reorderedNums := nums
	for {
		valid, i, j := isValidUpdate(reorderedNums, mustComeBefore)
		if valid {
			break
		}
		// Swap the two offending numbers and try again.
		reorderedNums[i], reorderedNums[j] = reorderedNums[j], reorderedNums[i]
	}
	return reorderedNums
}
