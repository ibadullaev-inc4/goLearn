package main

import (
	"fmt"
	"os"
	"quiz/game"
	"quiz/questions"
	"quiz/suffler"
)

func main() {
	questions, err := questions.LoadQuestions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't load questions %v\n", err)
		os.Exit(1)
	}

	suffler.Suffle(questions)
	correctAnswers := game.Run(questions)
	fmt.Printf("Correct answers: %d", correctAnswers)

}
