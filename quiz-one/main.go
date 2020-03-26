package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	//open file
	fileName := flag.String("file", "problems.csv", "a file name")
	waitTimeInSeconds := flag.Int("waitTime", 2, "this is the time in seconds to wait for an answer to a quiz question " +
		"from a user")
	flag.Parse()

	csvFile, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("Couldn't open the csv file ", err)
	}

	//Parse the csv file
	r := csv.NewReader(csvFile)
	score := 0
	total := 0
	reader := bufio.NewReader(os.Stdin)
	for {
		line , err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		total++
		fmt.Printf("Problem %d: %s = ",total, line[0]) //prints the question to the user

		answerTimer := time.NewTimer( time.Duration(*waitTimeInSeconds) * time.Second)

		go func() {
			<-answerTimer.C
			printQuizResult(score,total)
			os.Exit(1)
		}()

		answer, _ := reader.ReadString('\n')
		answerTimer.Stop()

		userAnswer := strings.TrimSpace(answer)
		correctAnswer := strings.TrimSpace(line[1])

		if userAnswer == correctAnswer {
			score++
		}
	}

	printQuizResult(score,total)
}

func printQuizResult(score, totalQuestions int)  {
	fmt.Printf("You scored %d out of %d questions\n", score, totalQuestions)
}