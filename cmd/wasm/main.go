package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	white      bool
	lastToggle time.Time
}

func (g *Game) Update() error {
	if time.Since(g.lastToggle) > time.Second {
		g.white = !g.white
		g.lastToggle = time.Now()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	c := color.Black
	if g.white {
		c = color.White
	}
	screen.Fill(c)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Square toggle")

	game := &Game{white: true, lastToggle: time.Now()}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
