package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/michaelmcallister/conway/life"
)

const width, height = 320, 240

// Game satisfies the ebiten.Game interface.
type Game struct {
	life *life.Life
}

// Update advances Life by one iteration.
func (g *Game) Update(screen *ebiten.Image) error {
	g.life.Step()
	return nil
}

// Draw takes the current board state and sets each alive cell to 0xFF (black).
func (g *Game) Draw(screen *ebiten.Image) {
	pix := make([]byte, width*height*4)
	for i, v := range g.life.BoardState() {
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
	}
	screen.ReplacePixels(pix)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func main() {
	life := life.New(width, height)

	ebiten.SetWindowTitle("Conways Game Of Life")
	game := &Game{life: life}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
