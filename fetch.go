package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// FetchInput downloads the input for a given day from adventofcode.com
func FetchInput(day string, year string) error {
	// Get session cookie from environment
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return fmt.Errorf("AOC_SESSION environment variable not set\n\nTo get your session:\n1. Log in to adventofcode.com\n2. Open browser dev tools (F12)\n3. Go to Application/Storage > Cookies\n4. Copy the value of the 'session' cookie\n5. Set it: export AOC_SESSION=your_session_here")
	}

	// Construct URL
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Add session cookie
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch input: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("received status code %d (perhaps the puzzle isn't unlocked yet?)", resp.StatusCode)
	}

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	// Write to file
	dayFolder := fmt.Sprintf("day%02s", day)
	inputPath := filepath.Join(dayFolder, "input.txt")

	if err := os.WriteFile(inputPath, body, 0644); err != nil {
		return fmt.Errorf("failed to write input file: %w", err)
	}

	fmt.Printf("✓ Downloaded input for day %s to %s\n", day, inputPath)
	return nil
}

// SubmitAnswer submits an answer to adventofcode.com
func SubmitAnswer(day, part, answer, year string) error {
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return fmt.Errorf("AOC_SESSION environment variable not set")
	}

	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/answer", year, day)

	// Create form data
	formData := fmt.Sprintf("level=%s&answer=%s", part, answer)

	req, err := http.NewRequest("POST", url, strings.NewReader(formData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to submit answer: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	bodyStr := string(body)

	// Parse response (this is a simplified version - you'd want better parsing)
	if strings.Contains(bodyStr, "That's the right answer") {
		fmt.Println("✓ Correct! ⭐")
		return nil
	} else if strings.Contains(bodyStr, "That's not the right answer") {
		fmt.Println("✗ Incorrect answer")
		if strings.Contains(bodyStr, "too high") {
			fmt.Println("  (too high)")
		} else if strings.Contains(bodyStr, "too low") {
			fmt.Println("  (too low)")
		}
		return nil
	} else if strings.Contains(bodyStr, "You gave an answer too recently") {
		fmt.Println("⏱  Wait a bit before submitting again")
		return nil
	} else {
		fmt.Println("Unknown response - check adventofcode.com")
		return nil
	}
}
