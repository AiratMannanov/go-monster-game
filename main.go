package main

import (
	"github.com/monsterslayer/acitons"
	"github.com/monsterslayer/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := ""
	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerHealth int
	var monsterHealth int
	var playerAttackDmg int
	var playerHealValue int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackDmg = acitons.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = acitons.HealPlayer()
	} else {
		playerAttackDmg = acitons.AttackMonster(true)
	}

	monsterAttackDmg = acitons.AttackPlayer()
	playerHealth, monsterHealth = acitons.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerHealth:     playerHealth,
		MonsterHelath:    monsterHealth,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
	}

	roundData.PrintRoundStatisctic()

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
