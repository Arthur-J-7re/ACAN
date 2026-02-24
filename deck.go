package main

import "fmt"

func testShuffle(shuffles []string, startingId int) []Card {
	deck := createDeck()
	start := 51 - startingId
	fmt.Println(shuffles, startingId)

	for _, shuffle := range shuffles {
		deck, start = shuffleDeck(deck, shuffle, start)
		fmt.Println(start, deck[start])
		fmt.Println()
	}

	return deck
}

func shuffleDeck(deck []Card, shuffle string, start int) ([]Card, int) {
	deck1 := deck[:26]
	deck2 := deck[26:]

	if shuffle == "IN" {
		newId := 0
		if start < 26 {
			newId = (start)*2 + 1
		} else {
			newId = (start - 26) * 2
		}
		return fuse(deck2, deck1), newId
	} else {
		newId := 0
		if start < 26 {
			newId = start * 2
		} else {
			newId = (start-26)*2 + 1
		}
		return fuse(deck1, deck2), newId
	}
}

func fuse(deck1 []Card, deck2 []Card) []Card {
	fused := []Card{}
	for k := range deck1 {
		fused = append(fused, deck1[k])
		fused = append(fused, deck2[k])
	}
	return fused
}

func createDeck() []Card {
	deck := []Card{}
	for _, color := range couleurs {
		for _, value := range values {
			deck = append(deck, Card{value, color})
		}
	}
	return deck
}
