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
	csvFileName := flag.String("csv", "problems.csv", "a csv file format of questions and answer")
	timeLimit := flag.Int("limit", 30, "Timer for quiz in seconds")
	flag.Parse()
	//_ = csvFileName
	//_ = timeLimit

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file : %s", *csvFileName))

	}
	read := csv.NewReader(file)
	lines, err := read.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	problems := parseLine(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	//<-timer.C // waits until this channel got message
	fmt.Printf("Total Given Time is %d seconds. You should finish before that!!\n\n", *timeLimit)

	correct := 0
problemLoop:
	for index, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", index+1, problem.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			//fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			//return
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == problem.answer {
				fmt.Println("Correct!")
				correct++
			}

			/*default:
				fmt.Printf("Problem #%d: %s = ", index+1, problem.question)
				var answer string
				fmt.Scanf("%s\n", &answer)
				if answer == problem.answer {
					fmt.Println("Correct!")
					correct++
				}
			}*/
		}
	}

	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))

}

type problem struct {
	question string
	answer   string
}

func parseLine(lines [][]string) []problem {
	returnValue := make([]problem, len(lines))

	for i, line := range lines {
		returnValue[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return returnValue
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
