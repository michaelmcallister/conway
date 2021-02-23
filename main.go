package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/michaelmcallister/conway/life"
)

const width, height = 320, 240

// conway satisfies the ebiten.Game interface.
type conway struct {
	Life   *life.Life
	buffer []byte
}

// Update advances Life by one iteration.
func (c *conway) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		c.Life.Set(x, y, true)
		return nil
	}
	c.Life.Step()
	c.buffer = make([]byte, width*height*4)
	for i, ok := range c.Life.BoardState() {
		if ok {
			c.buffer[4*i] = 0xff
			c.buffer[4*i+1] = 0xff
			c.buffer[4*i+2] = 0xff
			c.buffer[4*i+3] = 0xff
		}
	}
	return nil
}

// Draw takes the current board state and sets each alive cell to 0xFF (black).
func (c *conway) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(c.buffer)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (*conway) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func main() {
	ebiten.SetWindowTitle("Conways Game Of Life")
	ebiten.SetMaxTPS(20)
	c := &conway{Life: life.New(width, height)}
	if err := ebiten.RunGame(c); err != nil {
		panic(err)
	}
}
