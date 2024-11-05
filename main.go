package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	listAllQuizzes()

	file, err := selectQuiz()
	if err != nil {
		log.Fatalf("CSV can't be read - %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error with reader")
	}

	correct := 0
	for i := 0; i < len(lines); i++ {
		var userAnswer string
		question := lines[i][0]
		answer := lines[i][1]
		fmt.Printf("Question %v: %v\n", i+1, question)
		fmt.Scanln(&userAnswer)
		if string(userAnswer) == answer {
			correct += 1
			fmt.Println("Great job, that was right!")
		} else {
			fmt.Println("Oof, not right, sorry")
		}
	}

	fmt.Printf("You got %v correct!", correct)

}

func selectQuiz() (*os.File, error) {
	var quizSelection string
	fmt.Println("Which quiz do you want to do?")
	fmt.Scanln(&quizSelection)
	quizSelection = "./quizzes/" + quizSelection + ".csv"
	return os.Open(quizSelection)
}

func listAllQuizzes() {
	quizList, err := os.ReadDir("./quizzes")
	if err != nil {
		log.Fatalf("Error reading quiz directory: %v", err)
	}
	fmt.Println(quizList)
}
