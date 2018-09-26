package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

/*
	Struct for csv line
*/
type CsvLine struct {
	Question string
	Answer   string
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	filename := "problems.csv"
	problems := readCsvFile(filename)
	randomNum := rand.Intn(len(problems))
	currentProblem := problems[randomNum]

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(currentProblem.Question)
	answer, err := reader.ReadString('\n')
	checkErr(err)
	currentAnswer := strings.Replace(answer, "\n", "", -1)
	if currentAnswer == currentProblem.Answer {
		fmt.Println("Hooray, you got it right!")
	} else {
		fmt.Println("Your answer is ", answer, "Good luck next time")
	}
}

func readCsvFile(fileName string) []CsvLine {
	var data []CsvLine
	f, err := os.Open(fileName)

	defer f.Close()

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	checkErr(err)

	for _, line := range lines {
		data = append(data, CsvLine{
			Question: line[0],
			Answer:   line[1],
		})
	}
	return data
}
