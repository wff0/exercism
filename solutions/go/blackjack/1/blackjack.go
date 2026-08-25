package blackjack

import "strconv"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "jack", "queen", "king", "ten":
		return 10
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		value, _ := strconv.Atoi(card)
		return value
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	if sum == 21 {
		if dealer < 10 {
			return "W"
		}
	}
	if sum >= 17 {
		return "S"
	}
	if sum >= 12 {
		if dealer < 7 {
			return "S"
		}
	}
	return "H"
}
