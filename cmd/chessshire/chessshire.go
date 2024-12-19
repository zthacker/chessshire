package main

import (
	"ChessShire/internal/factory"
	"ChessShire/internal/games/defaultgame"
	"ChessShire/internal/player"
)

func main() {

	playerOne := factory.NewDefaultPlayer()
	playerTwo := factory.NewDefaultPlayer()

	players := [2]*player.Player{}
	players[0] = playerOne
	players[1] = playerTwo

	defaultGame := defaultgame.NewGame(players)
	defaultGame.Play()

}
