package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "init":
		handleInit(os.Args[2:])
	case "fetch":
		handleFetch(os.Args[2:])
	case "test":
		handleTest(os.Args[2:])
	case "submit":
		handleSubmit(os.Args[2:])
	case "bench":
		handleBench(os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Advent of Code CLI Tool")
	fmt.Println("\nUsage:")
	fmt.Println("  aoc init <day>           Create folder structure for a day")
	fmt.Println("  aoc fetch <day>          Download input for a day")
	fmt.Println("  aoc test <day>           Run tests for a day")
	fmt.Println("  aoc submit <day> <part> <answer>  Submit an answer")
	fmt.Println("  aoc bench <day>          Run benchmarks for a day")
	fmt.Println("\nExamples:")
	fmt.Println("  aoc init 1")
	fmt.Println("  aoc fetch 1")
	fmt.Println("  aoc test 1")
	fmt.Println("  aoc submit 1 1 42")
}

func handleInit(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: day number required")
		fmt.Println("Usage: aoc init <day>")
		os.Exit(1)
	}
	day := args[0]

	if err := InitDay(day); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func handleFetch(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: day number required")
		fmt.Println("Usage: aoc fetch <day>")
		os.Exit(1)
	}
	day := args[0]

	year := os.Getenv("AOC_YEAR")
	if year == "" {
		year = "2024"
	}

	if err := FetchInput(day, year); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func handleTest(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: day number required")
		fmt.Println("Usage: aoc test <day>")
		os.Exit(1)
	}
	day := args[0]

	if err := RunTests(day); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func handleSubmit(args []string) {
	if len(args) < 3 {
		fmt.Println("Error: day, part, and answer required")
		fmt.Println("Usage: aoc submit <day> <part> <answer>")
		os.Exit(1)
	}
	day := args[0]
	part := args[1]
	answer := args[2]

	year := os.Getenv("AOC_YEAR")
	if year == "" {
		year = "2024"
	}

	if err := SubmitAnswer(day, part, answer, year); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func handleBench(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: day number required")
		fmt.Println("Usage: aoc bench <day>")
		os.Exit(1)
	}
	day := args[0]

	if err := RunBenchmarks(day); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
