package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	score := 0
	csvFilename := flag.String("csv", "questions.csv", "CSV file")
	timeLimit := flag.Int("limit", 30, "Time limit in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("ERROR")
		os.Exit(1)
	}
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		os.Exit(1)
	}

  	problemLoop:
	for _, record := range lines {
		fmt.Printf("%s\n", strings.Join(record[:len(record)-1], ","))

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("Time exceeded\n")
			break problemLoop
		case answer := <-answerCh:
			if answer == record[len(record)-1] {
				score += 1
			}
		}
	}

	fmt.Printf("You scored %v out of %v\n", score, len(lines))
}
