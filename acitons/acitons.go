package acitons

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) {
	minAttachValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttachValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}
	dmgValue := generateRandBetween(minAttachValue, maxAttackValue)
	currentMonsterHealth -= dmgValue
}

func HealPlayer() {
	minAttachValue := PLAYER_HEAL_MIN_VALUE
	maxAttackValue := PLAYER_HEAL_MAX_VALUE

	healValue := generateRandBetween(minAttachValue, maxAttackValue)

	healthDiff := PLAYER_HEALTH - currentMonsterHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
	}
}

func AttackPlayer() {
	minAttachValue := MONSTER_ATTACK_MIN_DMG
	maxAttackValue := MONSTER_ATTACK_MAX_DMG

	dmgValue := generateRandBetween(minAttachValue, maxAttackValue)
	currentPlayerHealth -= dmgValue
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
