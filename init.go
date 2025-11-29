package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// InitDay creates the folder structure and template files for a given day
func InitDay(day string) error {
	// Create day folder (e.g., day01, day02)
	dayFolder := fmt.Sprintf("day%02s", day)

	if err := os.MkdirAll(dayFolder, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create main.go file
	mainPath := filepath.Join(dayFolder, "main.go")
	if err := os.WriteFile(mainPath, []byte(mainTemplate(day)), 0644); err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}

	// Create test file
	testPath := filepath.Join(dayFolder, "main_test.go")
	if err := os.WriteFile(testPath, []byte(testTemplate(day)), 0644); err != nil {
		return fmt.Errorf("failed to create test file: %w", err)
	}

	// Create empty input file
	inputPath := filepath.Join(dayFolder, "input.txt")
	if err := os.WriteFile(inputPath, []byte(""), 0644); err != nil {
		return fmt.Errorf("failed to create input file: %w", err)
	}

	// Create example input file
	examplePath := filepath.Join(dayFolder, "example.txt")
	if err := os.WriteFile(examplePath, []byte(""), 0644); err != nil {
		return fmt.Errorf("failed to create example file: %w", err)
	}

	fmt.Printf("✓ Created %s/\n", dayFolder)
	fmt.Printf("  ├── main.go\n")
	fmt.Printf("  ├── main_test.go\n")
	fmt.Printf("  ├── input.txt\n")
	fmt.Printf("  └── example.txt\n")

	return nil
}

func mainTemplate(day string) string {
	return fmt.Sprintf(`package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read example and input
	example, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	exampleLines := strings.Split(strings.TrimSpace(string(example)), "\n")
	inputLines := strings.Split(strings.TrimSpace(string(input)), "\n")

	// Part 1 Example
	part1Example := solvePart1(exampleLines)
	fmt.Printf("Part 1 Example: %%d\n", part1Example)

	// Part 1 Input
	part1Input := solvePart1(inputLines)
	fmt.Printf("Part 1 Input: %%d\n", part1Input)

	// Part 2 Example
	part2Example := solvePart2(exampleLines)
	fmt.Printf("Part 2 Example: %%d\n", part2Example)

	// Part 2 Input
	part2Input := solvePart2(inputLines)
	fmt.Printf("Part 2 Input: %%d\n", part2Input)
}

func solvePart1(lines []string) int {
	// TODO: Implement part 1
	return 0
}

func solvePart2(lines []string) int {
	// TODO: Implement part 2
	return 0
}
`)
}

func testTemplate(day string) string {
	return `package main

import (
	"os"
	"strings"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input, err := os.ReadFile("example.txt")
	if err != nil {
		t.Fatalf("failed to read example: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	result := solvePart1(lines)
	
	expected := 0 // TODO: Update with expected value
	if result != expected {
		t.Errorf("Part 1 = %d; want %d", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("example.txt")
	if err != nil {
		t.Fatalf("failed to read example: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	result := solvePart2(lines)
	
	expected := 0 // TODO: Update with expected value
	if result != expected {
		t.Errorf("Part 2 = %d; want %d", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solvePart1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solvePart2(lines)
	}
}
`
}
