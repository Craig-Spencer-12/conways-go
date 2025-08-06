package main

import (
	"image/color"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenHeight = 2048
	screenWidth  = 2048
	cellSize     = 16

	numRows = screenHeight / cellSize
	numCols = screenWidth / cellSize
)

type Board = [numCols][numRows]bool

type Game struct {
	state  Board
	paused bool
	fps    time.Duration

	lastUpdate time.Time
}

func (g *Game) Update() error {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		col, row := x/cellSize, y/cellSize
		g.state[col][row] = !g.state[col][row]
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.paused = !g.paused
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.fps = 10
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.fps = 15
	} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
		g.fps = 30
	} else if inpututil.IsKeyJustPressed(ebiten.Key4) {
		g.fps = 60
	} else if inpututil.IsKeyJustPressed(ebiten.Key5) {
		g.fps = math.MaxInt
	}

	if g.paused {
		return nil
	}

	now := time.Now()
	if now.Sub(g.lastUpdate) < time.Second/g.fps {
		return nil
	}
	g.lastUpdate = now

	var newState Board
	for i, col := range g.state {
		for j := range col {
			newState[i][j] = g.aliveNextTick(i, j)
		}
	}

	g.state = newState
	return nil
}

// Conway's Game of Life Rules:
// 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// 2. Any live cell with two or three live neighbours lives on to the next generation.
// 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
// 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
func (g *Game) aliveNextTick(x, y int) bool {
	count := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i >= 0 && i < len(g.state) && j >= 0 && j < len(g.state[0]) && g.state[i][j] {
				count++
			}
		}
	}

	isAlive := g.state[x][y]
	if isAlive && count >= 3 && count <= 4 {
		return true
	} else if !isAlive && count == 3 {
		return true
	}

	return false
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.White)

	for x := 0; x <= screenWidth; x += cellSize {
		vector.StrokeLine(screen, float32(x), 0, float32(x), screenHeight, 1, color.Black, true)
	}

	for y := 0; y <= screenHeight; y += cellSize {
		vector.StrokeLine(screen, 0, float32(y), screenWidth, float32(y), 1, color.Black, true)
	}

	for i, col := range g.state {
		for j, cellIsAlive := range col {
			if cellIsAlive {
				vector.DrawFilledRect(screen, float32(i*cellSize), float32(j*cellSize), cellSize, cellSize, color.Black, true)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetTPS(ebiten.SyncWithFPS)
	ebiten.SetWindowSize(1024, 1024)
	ebiten.SetWindowTitle("Conway's Game of Life")

	game := &Game{paused: true, fps: 10, lastUpdate: time.Now()}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
