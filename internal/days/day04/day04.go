package day04

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkLine(line string) int {
	count := 0
	// Check how many times "XMAS" and "SAMX" occur in line
	for i := 0; i < len(line)-3; i++ {
		if line[i:i+4] == "XMAS" || line[i:i+4] == "SAMX" {
			count++
		}
	}
	return count
}

func SolvePart1(input string) (int, error) {
	lines, i, err := getInputLines(input)
	if err != nil {
		return i, err
	}
	total := 0
	for _, line := range lines {
		//fmt.Printf("Checking horizontal line: %s\n", line)
		total += checkLine(line)
	}
	// Check vertical columns
	for col := 0; col < len(lines[0]); col++ {
		var column string
		for row := 0; row < len(lines); row++ {
			if col < len(lines[row]) {
				column += string(lines[row][col])
			}
		}
		//fmt.Printf("Checking vertical column: %s\n", column)
		total += checkLine(column)
	}
	// Check diagonals
	// Direction (top-left to bottom-right) "\\"
	for row := 0; row < len(lines); row++ {
		var diagonal string
		for r, c := row, 0; r < len(lines) && c < len(lines[r]); r, c = r+1, c+1 {
			diagonal += string(lines[r][c])
		}
		if len(diagonal) >= 4 {
			//fmt.Printf("Checking diagonal (\\) starting at row %d: %s\n", row, diagonal)
			total += checkLine(diagonal)
		}
	}
	for col := 1; col < len(lines[0]); col++ { // start at top row, each col (except 0 already covered)
		var diagonal string
		for r, c := 0, col; r < len(lines) && c < len(lines[r]); r, c = r+1, c+1 {
			diagonal += string(lines[r][c])
		}
		if len(diagonal) >= 4 {
			//fmt.Printf("Checking diagonal (\\) starting at column %d: %s\n", col, diagonal)
			total += checkLine(diagonal)
		}
	}
	// Direction (top-right to bottom-left) "/" (missing previously)
	lastCol := len(lines[0]) - 1
	// Start from top row, every column
	for col := 0; col <= lastCol; col++ {
		var diagonal string
		for r, c := 0, col; r < len(lines) && c >= 0; r, c = r+1, c-1 {
			if c < len(lines[r]) { // guard for ragged lines
				diagonal += string(lines[r][c])
			} else {
				break
			}
		}
		if len(diagonal) >= 4 {
			//fmt.Printf("Checking diagonal (/) starting at top row col %d: %s\n", col, diagonal)
			total += checkLine(diagonal)
		}
	}
	// Start from left-most of bottom side (row 1 - end) at last column
	for row := 1; row < len(lines); row++ {
		var diagonal string
		for r, c := row, lastCol; r < len(lines) && c >= 0; r, c = r+1, c-1 {
			if c < len(lines[r]) { // guard for ragged lines
				diagonal += string(lines[r][c])
			} else {
				break
			}
		}
		if len(diagonal) >= 4 {
			//fmt.Printf("Checking diagonal (/) starting at row %d lastCol: %s\n", row, diagonal)
			total += checkLine(diagonal)
		}
	}
	if total == 0 {
		return 0, fmt.Errorf("no valid patterns found")
	}

	return total, nil
}

func getInputLines(input string) ([]string, int, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("failed to close input file: %v\n", err)
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
		return nil, 0, err
	}
	if len(lines) == 0 {
		return nil, 0, fmt.Errorf("no input data found")
	}
	return lines, 0, nil
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
	total := 0
	rows := len(lines)
	for r := 1; r < rows-1; r++ {
		cols := len(lines[r])
		for c := 1; c < cols-1; c++ {
			if lines[r][c] != 'A' {
				continue
			}
			// Ensure neighbor rows have enough columns (handle potential ragged lines gracefully)
			if c >= len(lines[r-1]) || c >= len(lines[r+1]) || c-1 >= len(lines[r-1]) || c+1 >= len(lines[r-1]) || c-1 >= len(lines[r+1]) || c+1 >= len(lines[r+1]) {
				continue
			}
			TL, TR := lines[r-1][c-1], lines[r-1][c+1]
			BL, BR := lines[r+1][c-1], lines[r+1][c+1]
			// Diagonal 1: TL A BR must be MAS or SAM
			diag1 := (TL == 'M' && BR == 'S') || (TL == 'S' && BR == 'M')
			// Diagonal 2: TR A BL must be MAS or SAM
			diag2 := (TR == 'M' && BL == 'S') || (TR == 'S' && BL == 'M')
			if diag1 && diag2 {
				total++
			}
		}
	}

	return total, nil
}
