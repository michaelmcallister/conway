package life

import (
	"math/rand"
	"time"
)

// Life represents the board state and internal counters.
type Life struct {
	width, height int
	Iteration     int
	board         []bool
}

// Positions contains the relative distance between the current
// cell and it's neighbours.
var positions = [8][2]int{
	{-1, 1}, {0, 1}, {1, 1},
	{-1, 0}, {1, 0},
	{-1, -1}, {0, -1}, {1, -1},
}

// New initializes a new world, containing a random initialized state.
func New(width, height int) *Life {
	l := &Life{width: width, height: height, board: make([]bool, width*height)}
	l.initRand()
	return l
}

// initRand will flip a coin and set the cell to alive based on the outcome.
func (l *Life) initRand() {
	rand.Seed(time.Now().UnixNano())
	for x := 0; x < l.width; x++ {
		for y := 0; y < l.height; y++ {
			alive := rand.Float32() > 0.5
			l.set(x, y, alive)
		}
	}
}

// Neighbours counts the amount of alive cells adjacent to the x,y co-ordinate.
func (l *Life) neighbours(x, y int) int {
	count := 0
	for _, pos := range positions {
		if l.alive(x+pos[0], y+pos[1]) {
			count++
		}
	}
	return count
}

func (l *Life) alive(x, y int) bool {
	x += l.width
	x %= l.width
	y += l.height
	y %= l.height
	return l.board[y*l.width+x]
}

func (l *Life) set(x, y int, alive bool) {
	x += l.width
	x %= l.width
	y += l.height
	y %= l.height
	l.board[y*l.width+x] = alive
}

// BoardState returns the current state as a slice of bools.
func (l *Life) BoardState() []bool {
	return l.board
}

// Step advances the board state applying Conways rules.
// 1. Any live cell with two or three live neighbours survives.
// 2. Any dead cell with three live neighbours becomes a live cell.
// 3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.
func (l *Life) Step() {
	buff := make([]bool, l.width*l.height)
	for y := 0; y < l.height; y++ {
		for x := 0; x < l.width; x++ {
			n := l.neighbours(x, y)
			if (n == 2 && l.alive(x, y)) || n == 3 {
				buff[y*l.width+x] = true
			}
		}
	}
	l.Iteration++
	l.board = buff
}
