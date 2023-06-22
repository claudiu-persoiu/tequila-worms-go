package game

import "math/rand"

var messages = []string{
	" bites the dust",
	" killed himself",
	" fallowed bin laden",
	", die die!",
	" didn't have 9 lives",
	" see you in hell",
	" has been undefined",
	" sleeps with the fish",
	" equals null",
}

func randomDeadMessage() string {
	return messages[rand.Intn(len(messages))]
}
