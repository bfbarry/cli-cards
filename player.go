package main

import (
	"fmt"
	"net"
)

type Pid int
type Player struct {
    id Pid
    hand []Card
	conn net.Conn
}

func (self *Player) deal(game *GameState, num int) {
	for range num {
		deckSize := len(game.deck)
		picked   := game.deck[deckSize-1]
		game.deck = game.deck[:deckSize-1]
		self.hand = append(self.hand, picked)
	}
}

func (self *Player) showCards() {
	out := ""
	for _, v := range self.hand {
		out += fmt.Sprintf("|%s%s| ", v.num, v.suit)
	}
	out += "\n"

	_, err := fmt.Fprintf(self.conn, "%s", out)
	if err != nil {
		fmt.Println("Write error:", err)
	}
}