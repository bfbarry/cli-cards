package main

import (
	"fmt"
	"net"
)


func (self *Player) deal(game *GameState, num int) {
	for range num {
		deckSize := len(game.deck)
		picked   := game.deck[deckSize-1]
		game.deck = game.deck[:deckSize-1]
		self.hand = append(self.hand, picked)
	}
}

func (self *Player) showCards(conn net.Conn) {
	out := ""
	for _, v := range self.hand {
		out += fmt.Sprintf("%s%s ", v.suit, v.num)
	}

	_, err := fmt.Fprintf(conn, "%s", out)
	if err != nil {
		fmt.Println("Write error:", err)
	}
}