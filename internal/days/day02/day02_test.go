package day02

import (
	"path/filepath"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	// Test with the actual input file
	inputPath := filepath.Join("input.txt")
	result, err := SolvePart1(inputPath)
	if err != nil {
		t.Fatalf("SolvePart1() error = %v", err)
	}

	expected := 660
	if result != expected {
		t.Errorf("SolvePart1() = %d, want %d", result, expected)
	}
}

func TestSolvePart2(t *testing.T) {
	// Test with the actual input file
	inputPath := filepath.Join("input.txt")
	result, err := SolvePart2(inputPath)
	if err != nil {
		t.Fatalf("SolvePart2() error = %v", err)
	}

	expected := 18934359
	if result != expected {
		t.Errorf("SolvePart2() = %d, want %d", result, expected)
	}
}
