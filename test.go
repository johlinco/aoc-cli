package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunTests executes the tests for a given day
func RunTests(day string) error {
	dayFolder := fmt.Sprintf("day%02s", day)

	// Check if folder exists
	if _, err := os.Stat(dayFolder); os.IsNotExist(err) {
		return fmt.Errorf("day %s doesn't exist - run 'aoc init %s' first", day, day)
	}

	fmt.Printf("Running tests for day %s...\n\n", day)

	cmd := exec.Command("go", "test", "-v")
	cmd.Dir = dayFolder
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("tests failed: %w", err)
	}

	return nil
}

// RunBenchmarks executes the benchmarks for a given day
func RunBenchmarks(day string) error {
	dayFolder := fmt.Sprintf("day%02s", day)

	// Check if folder exists
	if _, err := os.Stat(dayFolder); os.IsNotExist(err) {
		return fmt.Errorf("day %s doesn't exist - run 'aoc init %s' first", day, day)
	}

	fmt.Printf("Running benchmarks for day %s...\n\n", day)

	cmd := exec.Command("go", "test", "-bench=.", "-benchmem")
	cmd.Dir = dayFolder
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("benchmarks failed: %w", err)
	}

	return nil
}
