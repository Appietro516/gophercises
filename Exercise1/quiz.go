package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	score := 0
	file, err := os.Open("questions.csv")
	if err != nil {
		fmt.Printf("ERROR")
	}
	reader := csv.NewReader(file)

	//loop over file
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		//write question
		fmt.Printf("%s\n", strings.Join(record[:len(record)-1], ","))

		//read answer
		var answer string
		fmt.Scanln(&answer)
		if answer == record[len(record)-1] {
			score += 1
		} else {
			break
		}
	}
	file.Seek(0, 0)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("ERROR")
	}

	fmt.Printf("You scored %v out of %v\n", score, len(lines))
}
