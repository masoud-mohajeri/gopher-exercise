package quizGame

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func Quiz() {

	csvFilename := flag.String("csv", "./exercises/quizGame/problem.csv", "a csv file containing the questions in for,at of question,answer")
	timeLimit := flag.Int("limit", 5, "time limit")

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file: %s\n", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("failed ro pars provided csv file %s", err))
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	problems := parsLines(lines)
	for i, p := range problems {
		fmt.Printf("Problem %d: %s : \n", i+1, p.q)
		fmt.Println("===== Debugger 1")

		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Println("===== Debugger chan")
			fmt.Scanf("%s\n", &answer)
			fmt.Println("===== Debugger after input")
			answerCh <- answer
		}()
		fmt.Println("===== Debugger 2")

		select {
		case <-timer.C:
			fmt.Printf("Your score %d percent", correct*100/len(problems))
			return
		case answer := <-answerCh:

			if answer == p.a {
				fmt.Println("Correct!")
				correct++
			}
		}
		fmt.Println("===== Debugger after select")

	}

	fmt.Printf("--- Your score %d percent", correct*100/len(problems))

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parsLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
