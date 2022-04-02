package main

import (
	"embed"
	"game/engine"
	_ "image/png"
	"log"

	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREENWIDTH  int = 490
	SCREENHEIGHT int = 300
)

//go:embed resource
var images embed.FS

func main() {
	ebiten.SetWindowSize(SCREENWIDTH*2, SCREENHEIGHT*2)
	ebiten.SetWindowTitle("golang Dibalo2")
	gameStart := engine.NewGame(&images)
	gameStart.StartEngine()
	if err := ebiten.RunGame(gameStart); err != nil {
		log.Fatal(err)
	}
}
