package main

type deck struct {
	ID       string `json:"deck_id"`
	Shuffled string `json:"shuffled"`
	Cards    []string
}

func New(shuffled bool, cards []string) {

}
