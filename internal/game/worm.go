package game

import "math/rand"

const minLength = 3
const initialLength = 7

type worm struct {
	uuid      string
	name      string
	color     string
	direction string
	length    int
	pieces    []piece
}

type piece struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func generateColor() string {
	letters := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	r := "#"
	for i := 0; i < 6; i++ {
		r += letters[rand.Intn(len(letters))]
	}
	return r
}

func NewWorm(uuid string, name string, position piece) *worm {
	return &worm{
		uuid:      uuid,
		name:      name,
		color:     generateColor(),
		pieces:    generatePieces(initialLength, position),
		direction: getRandomDir(),
		length:    minLength,
	}
}

func (w *worm) Step() {
	movement := getMovementByDirection(w.direction)

	w.pieces = append([]piece{{X: w.pieces[0].X + movement.X, Y: w.pieces[0].Y + movement.Y}}, w.pieces...)
	w.pieces = w.pieces[:len(w.pieces)-1]
}

func generatePieces(size int, position piece) []piece {
	var r []piece
	for i := 0; i < size; i++ {
		r = append(r, position)
	}
	return r
}
