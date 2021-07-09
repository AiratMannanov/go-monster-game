package interaction

import "fmt"

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(isSpecialRound bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if isSpecialRound {
		fmt.Println("(2) Special Attack!")
	}
}
