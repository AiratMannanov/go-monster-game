package interaction

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHelath    int
}

func PrintGreeting() {
	asciiFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	asciiFigure.Print()

	fmt.Println()
	fmt.Println("Starting a new game")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(isSpecialRound bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if isSpecialRound {
		fmt.Println("(3) Special Attack!")
	}
}

func (roundData *RoundData) PrintRoundStatisctic() {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage. \n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed %v damage.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage. \n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health %v . \n", roundData.PlayerHealth)
	fmt.Printf("Monster Health %v . \n", roundData.MonsterHelath)
}

func DeclareWinner(winner string) {
	asciiFigure := figure.NewColorFigure("Game Over!", "", "red", true)
	asciiFigure.Print()
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {
	file, err := os.Create("gamelog.txt")
	if err != nil {
		fmt.Println("Saving a log file failed. Exiting")
		return
	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                fmt.Sprint(index + 1),
			"Action":               value.Action,
			"Player Attack Damage": fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":    fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Value": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health":        fmt.Sprint(value.PlayerHealth),
			"Monster Health":       fmt.Sprint(value.MonsterHelath),
		}

		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing into log file failed. Exiting")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote Data to log!")
}
