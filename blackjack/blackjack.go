package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
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
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

func ParseHand(card1, card2 string) int {
	return ParseCard(card1) + ParseCard(card2)
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseHand(card1, card2)
	switch {
	case card1 == card2 && card1 == "ace":
		return "P"
	case handScore == 21 && ParseCard(dealerCard) < 10:
		return "W"
	case
		handScore == 21 && ParseCard(dealerCard) >= 10,
		handScore >= 17 && handScore <= 20,
		handScore >= 12 && handScore <= 16 && ParseCard(dealerCard) < 7:

		return "S"

	case
		handScore <= 11,
		handScore >= 12 && handScore <= 16 && ParseCard(dealerCard) >= 7:

		return "H"
	default:
		panic("No Strategy")
	}
}
