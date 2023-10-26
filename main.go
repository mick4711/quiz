package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	// read problems from CSV file
	records := readCSVFile()

	// construct a slice of quiz of type Problem, structs with questions and answers
	quiz := makeQuiz(records)

	// display problems
	for i, problem := range quiz {
		fmt.Println("Problem:", i+1, "\tQ:", problem.question, "\tA:", problem.answer)
	}

	// TODO write tests
	// TODO add a flag for the filename
	// TODO handle missing / invalid file and invalid records
	// TODO create a slice for results with question number and correct /wrong flag
	// TODO write a question to stdout and take in response
	// TODO update results
	// TODO print out results

}

func readCSVFile() [][]string {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func makeQuiz(records [][]string) []Problem {
	var quiz []Problem
	for _, record := range records {
		var problem Problem
		problem.question = record[0]
		problem.answer = record[1]
		quiz = append(quiz, problem)
	}
	return quiz
}
