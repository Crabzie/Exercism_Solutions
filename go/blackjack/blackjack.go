// Package blackjack is used to calculate the result of first round of blackjack
package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
/*
   ace		11	eight	8
   two		2	  nine	9
   three	3	  ten		10
   four	  4	  jack	10
   five	  5	  queen	10
   six		6	  king	10
   seven	7	  other	0
*/
func ParseCard(card string) int {
	var value int
	switch card {
	default:
		value = 0
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten":
		value = 10
	case "jack":
		value = 10
	case "queen":
		value = 10
	case "king":
		value = 10
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
// dealer variable holds the integer value of the dealerCard variable
// card1Num variable holds the integer value of the card1 variable
// card2Num variable holds the integer value of the card2 variable
// sumResult variable holds the integer sum of card1 and card2
func FirstTurn(card1, card2, dealerCard string) string {
	dealer := ParseCard(dealerCard)
	card1Num := ParseCard(card1)
	card2Num := ParseCard(card2)
	sumResult := card1Num + card2Num
	var result string
	switch {
	case card1Num == 11 && card2Num == 11:
		result = "P"

	case (sumResult == 21) && (dealer != 10 && dealer != 11):
		result = "W"

	case (sumResult == 21) && (dealer == 10 || dealer == 11):
		result = "S"

	case (sumResult <= 20) && (sumResult >= 17):
		result = "S"

	case ((sumResult <= 16) && (sumResult >= 12)) && (dealer >= 7):
		result = "H"

	case ((sumResult <= 16) && (sumResult >= 12)) && (dealer < 7):
		result = "S"

	case sumResult <= 11:
		result = "H"

	}
	return result
}
