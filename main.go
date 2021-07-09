package main

import "fmt"

func main() {
	startGame()

	winner := ""
	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {

}

func executeRound() string {

}

func endGame() {

}
