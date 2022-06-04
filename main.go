package main

import (
	game "github.com/kapilyad/tennis-game/game"
	player "github.com/kapilyad/tennis-game/player"
)

func main() {
	p1 := player.NewPlayer("A", true)
	p2 := player.NewPlayer("B", false)
	game := game.NewGame(*p1, *p2)
	game.StartGame()
}
