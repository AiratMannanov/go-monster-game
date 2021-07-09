package main

import (
	"github.com/monsterslayer/interaction"
)

func main() {
	startGame()

	winner := ""
	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {

}

func endGame() {

}
