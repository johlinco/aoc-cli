 # AOC CLI

  A command-line tool for streamlining your Advent of Code workflow. Automatically
  generates project structure, downloads input data, runs tests, and submits solutions.

  ## Features

  - 🏗️ **Project Setup**: Auto-generate folder structure with templates for each day
  - 📥 **Input Fetching**: Download puzzle input directly from Advent of Code
  - 🧪 **Testing**: Run tests with example data
  - 📊 **Benchmarking**: Performance testing for your solutions
  - 🚀 **Submission**: Submit answers directly from the CLI

  ## Installation

  ### Prerequisites
  - Go 1.19 or later
  - Advent of Code session cookie (for fetching inputs and submitting)

  ### Build from Source
  ```bash
  git clone <repository-url>
  cd aoc-cli
  go build -o aoc

  Add to PATH (optional)

  # Move binary to a directory in your PATH
  sudo mv aoc /usr/local/bin/

  Configuration

  Session Cookie Setup

  To fetch inputs and submit answers, you need your Advent of Code session cookie:

  1. Login to https://adventofcode.com
  2. Open browser developer tools (F12)
  3. Go to Application/Storage → Cookies → adventofcode.com
  4. Copy the session value
  5. Set environment variable:

  export AOC_SESSION="your_session_cookie_here"

  Year Configuration

  Set the year (defaults to 2024):
  export AOC_YEAR="2024"

  Usage

  Initialize a New Day

  Creates folder structure with template files:
  ./aoc init <day>

  This creates:
  day01/
  ├── main.go       # Solution template with Part 1 & 2 functions
  ├── main_test.go  # Test template with benchmarks
  ├── input.txt     # Empty file for puzzle input
  └── example.txt   # Empty file for example data

  Fetch Puzzle Input

  Downloads the official input for a day:
  ./aoc fetch <day>

  Run Tests

  Execute tests for a specific day:
  ./aoc test <day>

  Submit Solutions

  Submit your answer for a specific part:
  ./aoc submit <day> <part> <answer>

  Run Benchmarks

  Performance test your solutions:
  ./aoc bench <day>

  Workflow Example

  Complete workflow for Day 1:

  # 1. Create project structure
  ./aoc init 1

  # 2. Download puzzle input
  ./aoc fetch 1

  # 3. Edit day01/example.txt with the example from the problem
  # 4. Edit day01/main.go to implement your solution
  # 5. Update expected values in day01/main_test.go

  # 6. Test your solution
  ./aoc test 1

  # 7. Submit Part 1
  ./aoc submit 1 1 42

  # 8. Update solution for Part 2 and submit
  ./aoc submit 1 2 1337

  # 9. Check performance
  ./aoc bench 1

  Generated Code Structure

  The init command creates a standardized template:

  main.go

  package main

  import (
      "fmt"
      "os"
      "strings"
  )

  func main() {
      // Reads both example.txt and input.txt
      // Runs both parts on both datasets
      // Displays results in organized format
  }

  func solvePart1(lines []string) int {
      // TODO: Implement part 1
      return 0
  }

  func solvePart2(lines []string) int {
      // TODO: Implement part 2
      return 0
  }

  main_test.go

  Includes:
  - Unit tests for both parts using example data
  - Benchmarks for performance testing
  - Automatic file reading and parsing

  Tips

  - Add example data to example.txt from the problem description
  - Update expected test values after implementing your solution
  - Use the benchmarks to optimize performance for larger inputs
  - The tool automatically handles year selection via AOC_YEAR environment variable

  Troubleshooting

  "Session required" error

  Make sure your AOC_SESSION environment variable is set correctly.

  "File not found" errors

  Ensure you're running commands from the directory containing the day folders.

  Build issues

  Verify you have Go 1.19+ installed: go version

  Contributing

  Feel free to open issues or submit pull requests for improvements!
