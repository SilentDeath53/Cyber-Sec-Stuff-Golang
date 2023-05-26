package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func analyzeLog(logEntry string) {
	// Adjust the regular expression pattern according to your log format
	pattern := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}).*?(error|warning)`)
	matches := pattern.FindStringSubmatch(logEntry)

	if len(matches) > 0 {
		timestamp, err := time.Parse("2006-01-02 15:04:05", matches[1])
		if err != nil {
			log.Println("Error parsing timestamp:", err)
		}

		logType := matches[2]

		// You can customize the logic according to your requirements
		if logType == "error" {
			fmt.Println("Abnormal activity detected - Error log at:", timestamp)
		} else if logType == "warning" {
			fmt.Println("Potential attack attempt detected - Warning log at:", timestamp)
		}
	}
}

func main() {
	file, err := os.Open("application.log") // Replace "application.log" with the path to your log file
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logEntry := strings.TrimSpace(scanner.Text())
		analyzeLog(logEntry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning log file:", err)
	}
}
