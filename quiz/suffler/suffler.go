package suffler

import (
	"math/rand"
	"quiz/questions"
)

func Suffle(questions []questions.Question) {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}
