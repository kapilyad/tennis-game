package game

import (
	"fmt"
	"time"

	player "github.com/kapilyad/tennis-game/player"
)

type Game struct {
	Players       []player.Player
	HittingPlayer player.Player
	WaitingPlayer player.Player
	IsBallDropped bool
}

var (
	hittingPlayer player.Player
)

func (g *Game) HitBall(ballDropped chan string) int {
	if g.IsBallDropped {
		ballDropped <- "Ball is dropped"
		return 1
	}
	hittingPlayer = g.HittingPlayer
	fmt.Printf("Player %v Hits \n", g.HittingPlayer.Name)

	// change the state of the player
	g.HittingPlayer.State.IsHitting = false
	time.Sleep(time.Second)

	// change the state of game
	g.HittingPlayer = g.WaitingPlayer
	g.WaitingPlayer = hittingPlayer

	return g.HitBall(ballDropped)
}

func NewGame(p1 player.Player, p2 player.Player) *Game {
	var hittingPlayer, waitingPlayer player.Player
	if p1.IsServingFirst {
		hittingPlayer = p1
		waitingPlayer = p2
	} else {
		hittingPlayer = p2
		waitingPlayer = p1
	}
	return &Game{
		Players:       []player.Player{p1, p2},
		HittingPlayer: hittingPlayer,
		WaitingPlayer: waitingPlayer,
	}
}

func (g *Game) StartGame() {
	ballDropped := make(chan string)
	go g.HitBall(ballDropped)

	select {
	case isBallDropped := <-ballDropped:
		fmt.Println(isBallDropped)
	}
}
