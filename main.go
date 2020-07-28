package main

import (
	"sync"

	"github.com/hajimehoshi/ebiten"
	"github.com/michaelmcallister/conway/life"
)

var waitGroup sync.WaitGroup

const width, height = 320, 240

// conway satisfies the ebiten.Game interface.
type conway struct {
	*life.Life
}

// Update advances Life by one iteration.
func (c *conway) Update(screen *ebiten.Image) error {
	c.Step()
	return nil
}

// Draw takes the current board state and sets each alive cell to 0xFF (black).
func (c *conway) Draw(screen *ebiten.Image) {
	pix := make([]byte, width*height*4)
	for i, v := range c.BoardState() {
		waitGroup.Add(1)
		go func(i int, v bool) {
			defer waitGroup.Done()
			if v {
				pix[4*i] = 0
				pix[4*i+1] = 0
				pix[4*i+2] = 0
				pix[4*i+3] = 0
			} else {
				pix[4*i] = 0xff
				pix[4*i+1] = 0xff
				pix[4*i+2] = 0xff
				pix[4*i+3] = 0xff
			}
		}(i, v)
	}
	waitGroup.Wait()
	screen.ReplacePixels(pix)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (*conway) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func main() {
	ebiten.SetWindowTitle("Conways Game Of Life")
	game := &conway{life.New(width, height)}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
