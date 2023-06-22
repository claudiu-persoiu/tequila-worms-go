package game

import "math/rand"

type worm struct {
	uuid   string
	name   string
	color  string
	pieces []piece
}

type piece struct {
	x int
	y int
}

func generateColor() string {
	letters := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	r := ""
	for i := 0; i < 6; i++ {
		r += letters[rand.Intn(len(letters))]
	}
	return r
}

func NewWorm(uuid string, name string) *worm {
	return &worm{
		uuid:   uuid,
		name:   name,
		color:  generateColor(),
		pieces: nil,
	}
}

func getRandomWithBezel(x int, bezel int) int {
	return rand.Intn(x-(bezel*2)) + bezel
}
