package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Card struct {
	Color string `json:"color"`
	Value string `json:"value"`
}

var (
	couleurs = []string{
		"Pique", "Coeur", "Trefle", "Carreau",
	}
	values = []string{
		"As", "2", "3", "4", "5", "6", "7", "8",
		"9", "10", "Valet", "Dame", "Roi",
	}
)

func main() {
	http.HandleFunc("/api/shuffle", shuffleHandler)

	println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func shuffleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	value := query.Get("value")
	color := query.Get("color")
	placeStr := query.Get("place")

	// Validation
	if !includes(values, value) || !includes(couleurs, color) {
		http.Error(w, "invalid card", http.StatusBadRequest)
		return
	}

	place, err := strconv.Atoi(placeStr)
	if err != nil || place < 1 || place > 52 {
		http.Error(w, "invalid place", http.StatusBadRequest)
		return
	}

	card := Card{
		Color: color,
		Value: value,
	}

	// 🔥 Tes fonctions existantes
	shuffle, start, bin, euc := acan(card, place)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]any{shuffle, start, bin, euc})
}

func includes(list []string, mot string) bool {
	for _, elem := range list {
		if elem == mot {
			return true
		}
	}
	return false
}
