package game

import (
	"math/rand"
)

type direction struct {
	X int
	Y int
}

var directions = map[string]direction{
	"left":  {X: -1, Y: 0},
	"right": {X: 1, Y: 0},
	"up":    {X: 0, Y: 1},
	"down":  {X: 0, Y: -1},
}

var possibleDirs = []string{
	"left", "right", "up", "down",
}

func getMovementByDirection(dir string) direction {
	return directions[dir]
}

func getRandomDir() string {
	return possibleDirs[rand.Intn(len(possibleDirs))]
}

func isValidDirection(dir string) bool {
	for _, p := range possibleDirs {
		if p == dir {
			return true
		}
	}
	return false
}

func isReversed(oldDir, newDir string) bool {
	return directions[oldDir].X+directions[newDir].X == 0 && directions[oldDir].Y+directions[newDir].X == 0
}
