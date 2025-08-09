# Advent of Code 2024

This repository contains my solutions for [Advent of Code 2024](https://adventofcode.com/2024) written in Go.

## Project Structure

```
advent-of-code-2024/
├── cmd/aoc/main.go          # Main entry point
├── internal/
│   ├── days/                # Solutions for each day
│   │   ├── day01/
│   │   │   ├── day01.go     # Solution implementation
│   │   │   └── input.txt    # Puzzle input
│   │   ├── day02/
│   │   │   ├── day02.go
│   │   │   └── input.txt
│   │   └── ...              # Days 3-25
│   └── utils/
│       └── utils.go         # Common utility functions
├── bin/                     # Built binaries
└── go.mod
```

## Usage

### Running Solutions with Makefile (Recommended)

The project includes a Makefile for easier command execution:

```bash
# Run a specific day and part
make run DAY=1 PART=1
make run DAY=2 PART=2

# Quick shortcuts for common days
make day1    # Runs both part 1 and part 2 of day 1
make day2    # Runs both part 1 and part 2 of day 2

# Run all implemented solutions
make all

# Build the project
make build

# Run tests
make test

# Format code
make fmt

# Clean build artifacts
make clean

# Show all available targets
make help
```

### Running Solutions Directly with Go

You can also run solutions using command line flags directly:

```bash
# Run a specific day and part
go run ./cmd/aoc/main.go -day=1 -part=1
go run ./cmd/aoc/main.go -day=1 -part=2

# Run all implemented solutions
go run ./cmd/aoc/main.go -all

# Default behavior (day 1, part 1)
go run ./cmd/aoc/main.go
```

### Building

```bash
# Build the binary (using Makefile)
make build

# Or build directly with Go
go build -o bin/aoc ./cmd/aoc/main.go

# Run the built binary
./bin/aoc -day=1 -part=1
```

### Adding New Day Solutions

1. **Copy your input**: Paste your puzzle input into `internal/days/dayXX/input.txt`

2. **Implement the solution**: Edit `internal/days/dayXX/dayXX.go` and implement the `SolvePart1` and `SolvePart2` functions

3. **Add to main.go**: Add the import and entry to the solutions map in `cmd/aoc/main.go`:

```go
import (
    // ... existing imports
    "github.com/neilsmahajan/advent-of-code-2024/internal/days/dayXX"
)

var solutions = map[int]Solution{
    // ... existing solutions
    XX: {dayXX.SolvePart1, dayXX.SolvePart2},
}
```

4. **Test your solution**:

```bash
go run ./cmd/aoc/main.go -day=XX -part=1
go run ./cmd/aoc/main.go -day=XX -part=2
```

### Utility Functions

The `internal/utils` package provides common functions:

- `ReadLines(filename)` - Read all lines from a file
- `ReadInts(filename)` - Read integers from a file (one per line)
- `SplitInts(s, delimiter)` - Split string and convert to integers
- `AbsInt(x)` - Absolute value
- `MinInt(a, b)`, `MaxInt(a, b)` - Min/Max functions
- `SumInts(nums)` - Sum of integer slice
- `ParseGrid(lines)` - Parse 2D character grid
- `InBounds(grid, row, col)` - Check grid bounds

Example usage:

```go
import "github.com/neilsmahajan/advent-of-code-2024/internal/utils"

func SolvePart1(input string) (int, error) {
    lines, err := utils.ReadLines(input)
    if err != nil {
        return 0, err
    }

    // Process lines...
    return result, nil
}
```

## Development

### Running Tests

```bash
go test ./...
```

### Project Setup

The project structure is already set up with template files for all 25 days. Each day has a placeholder implementation that you can fill in with your solution.

## Solutions

- [x] Day 1: Historian Hysteria ⭐⭐
- [x] Day 2: Red-Nosed Reports ⭐⭐
- [ ] Day 3: TBD
- [ ] Day 4: TBD
- [ ] Day 5: TBD
- [ ] Day 6: TBD
- [ ] Day 7: TBD
- [ ] Day 8: TBD
- [ ] Day 9: TBD
- [ ] Day 10: TBD
- [ ] Day 11: TBD
- [ ] Day 12: TBD
- [ ] Day 13: TBD
- [ ] Day 14: TBD
- [ ] Day 15: TBD
- [ ] Day 16: TBD
- [ ] Day 17: TBD
- [ ] Day 18: TBD
- [ ] Day 19: TBD
- [ ] Day 20: TBD
- [ ] Day 21: TBD
- [ ] Day 22: TBD
- [ ] Day 23: TBD
- [ ] Day 24: TBD
- [ ] Day 25: TBD
