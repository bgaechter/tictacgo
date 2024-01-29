package main

import (
	"log"

	"github.com/bgaechter/tictacgo/tictacgo"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := tictacgo.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("TicTacGo")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
