package main

// A Deck represents a stack of cards from a specific cardset.
type deck struct {
	ID        string `json:"deck_id"`
	Shuffled  string `json:"shuffled"`
	Remaining string `json:"remaining"`
	CardsetID string `json:"cardset_id"`
	Cards     []card `json:"cards"`
}

func New(shuffled bool, cards []card) {

}

func (d *deck) Draw(quantity uint32) {

}
