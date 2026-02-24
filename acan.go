package main

import (
	"strconv"
)

func acan(card Card, place int) ([]string, int) {
	start := getStart(card)
	target := place - 1

	M := start * 64
	S := M - target

	X, Y := getBinaries(S)
	return getShuffle(X, Y), start
}

func getBinaries(S int) (string, string) {
	var q, r int
	if S%52 == 0 {
		q = S / 52
		r = 0
	} else {
		q = S/52 + 1
		r = 52*q - S
	}
	return getBinary(q), getBinary(r)
}

func getBinary(num int) string {
	bin := ""
	loop := num
	for loop > 0 {
		if loop%2 == 1 {
			bin = "1" + bin
			loop = (loop - 1) / 2
		} else {
			bin = "0" + bin
			loop = loop / 2
		}
	}
	for len(bin) < 6 {
		bin = "0" + bin
	}
	return bin
}

func getShuffle(X string, Y string) []string {
	shuffle := []string{}
	if len(X) != 6 || len(Y) != 6 {
	} else {
		for i := range X {
			if X[i] != Y[i] {
				shuffle = append(shuffle, "IN")
			} else {
				shuffle = append(shuffle, "OUT")
			}
		}
	}
	return shuffle
}

func getStart(card Card) int {
	order := getOrderOfCard(card)
	color := card.Color
	if order == -1 {
		return -1
	}
	switch color {
	case "Carreau":
		return order
	case "Trefle":
		return 13 + order
	case "Coeur":
		return 26 + order
	case "Pique":
		return 39 + order
	}
	return -1
}

func getOrderOfCard(card Card) int {
	switch card.Value {
	case "As":
		return 12
	case "Roi":
		return 0
	case "Dame":
		return 1
	case "Valet":
		return 2
	default:
		v, err := strconv.Atoi(card.Value)
		if err != nil {
			return -1 // valeur invalide
		}
		return 13 - v
	}
}
