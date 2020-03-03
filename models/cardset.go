package main

// A Cardset represents a game of card.
// It can be the typical 52 cards game, it could also be a preset for typical games requiring less cards or even a game of Pokemon cards with some upgrades to the Card model
type cardset struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Codes []string `json:"codes"`
}

func New(name string, codes []string) {

}
