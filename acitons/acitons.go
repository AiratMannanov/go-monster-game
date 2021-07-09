package acitons

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minAttachValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttachValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}
	dmgValue := generateRandBetween(minAttachValue, maxAttackValue)
	currentMonsterHealth -= dmgValue
	return dmgValue
}

func HealPlayer() int {
	minAttachValue := PLAYER_HEAL_MIN_VALUE
	maxAttackValue := PLAYER_HEAL_MAX_VALUE

	healValue := generateRandBetween(minAttachValue, maxAttackValue)

	healthDiff := PLAYER_HEALTH - currentMonsterHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
		return healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
		return healthDiff
	}

}

func AttackPlayer() int {
	minAttachValue := MONSTER_ATTACK_MIN_DMG
	maxAttackValue := MONSTER_ATTACK_MAX_DMG

	dmgValue := generateRandBetween(minAttachValue, maxAttackValue)
	currentPlayerHealth -= dmgValue
	return dmgValue
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

func generateRandBetween(min int, max int) int {
	// fmt.Println(min, max)
	return randGenerator.Intn(max-min) + min
}
