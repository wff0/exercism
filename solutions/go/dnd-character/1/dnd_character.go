package dndcharacter

import (
	"math"
	"math/rand/v2"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	tmp := math.Floor(float64(score-10) / 2)
	return int(tmp)
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	a := rand.IntN(6) + 1
	b := rand.IntN(6) + 1
	c := rand.IntN(6) + 1
	d := rand.IntN(6) + 1
	min := min(a, min(b, min(c, d)))
	return a + b + c + d - min
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	ch := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	ch.Hitpoints = 10 + Modifier(ch.Constitution)
	return ch
}
