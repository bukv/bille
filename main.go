package main

import (
	"bille/game"

	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(1000, 1000)

	game.Table(dc)

	game.BallMove(dc)

	dc.SavePNG("out/out.png")
}
