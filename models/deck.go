package main

// A Deck represents a stack of cards.
// At the moment it can contain cards from any cardset.
type deck struct {
	ID        string `json:"deck_id"`
	Shuffled  string `json:"shuffled"`
	Remaining string `json:"remaining"`
	Cards     []card `json:"cards"`
}

func New(shuffled bool, cards []card) {

}

func (d *deck) Draw(quantity uint32) {

}
