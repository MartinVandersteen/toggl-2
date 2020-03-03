package main

// A Card represents a game card and everything that makes it different from other cards.
// To accomodate for other games, new attributes could be added, like an image, an attack/defense value, etc.
type card struct {
	Code      string `json:"code"`
	Value     string `json:"value"`
	Suit      string `json:"suit"`
	CardsetID string `json:"cardset_id"`
}
