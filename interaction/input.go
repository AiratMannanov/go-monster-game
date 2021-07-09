package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(isSpecialRound bool) string {
	for {
		playerChoice, _ := getPlayrInput()
		if playerChoice == "1" {
			return "ATTACK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && isSpecialRound {
			return "SPECIAL_ATTACK"
		}
		fmt.Println("Fetching the user input failed. Please try again.")
	}
}

func getPlayrInput() (string, error) {
	fmt.Print("Your choice: ")
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput, nil
}
