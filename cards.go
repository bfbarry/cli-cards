package main

type Card struct {
    suit string
    num  string
}

func init_deck() []Card {
	cards := []Card{}
	suits := [4]string{"D", "C", "S", "H"}
	nums := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	var c Card;
	i := 0
	for _, s := range suits {
		for _, n := range nums {
			c = Card{suit: s, num: n}
			cards[i] = c
		}
	}

	return cards
}


