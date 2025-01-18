package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

func startOfNextNextQuarter() string {
	now := time.Now().UTC()
	quarter := (now.Month()-1)/3 + 1

	nextNextQuarter := (quarter + 2) % 4
	year := now.Year()

	if nextNextQuarter == 1 || nextNextQuarter == 2 {
		year++
	}

	if nextNextQuarter == 0 {
		nextNextQuarter = 4
	}

	// Determine the first day of the next next quarter
	startMonth := time.Month((nextNextQuarter-1)*3 + 1)
	startOfQuarter := time.Date(year, startMonth, 1, 0, 0, 0, 0, time.UTC)

	// Return the ISO 8601 formatted timestamp (RFC3339 ensures correctness)
	return startOfQuarter.Format(time.RFC3339)
}

// updateExpires reads, modifies, and writes back the file
func updateExpires(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	newExpires := fmt.Sprintf("Expires: %s", startOfNextNextQuarter())

	re := regexp.MustCompile(`Expires: .*Z`)
	updatedContent := re.ReplaceAllString(string(content), newExpires)

	return os.WriteFile(filePath, []byte(updatedContent), 0644)
}

func main() {
	filePath := "../../data/.well-known/security.txt"
	err := updateExpires(filePath)
	if err != nil {
		fmt.Println("Error updating file:", err)
	} else {
		fmt.Println("Successfully updated 'Expires' in", filePath)
	}
}
