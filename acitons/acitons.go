package acitons

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

func AttackMonster(isSpecialRound bool) {

}

func HealPlayer() {

}

func AttackPlayer() {

}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
