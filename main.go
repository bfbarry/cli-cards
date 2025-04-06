package main

import (
	"fmt"
	"math/rand"
	"net"
)


func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting TCP server:", err)
        return
    }
    defer ln.Close()

    fmt.Println("TCP server listening on :8080")

    init_players := make(map[Pid]Player)
    deck := initDeck()
    rand.Shuffle(len(deck), func(i, j int) {
        deck[i], deck[j] = deck[j], deck[i]
    })

    game := GameState{players: init_players, state: 0, deck: deck}

    for {
        p, err := game.createPlayer(ln)
        if err != nil {
            continue
        }

        go game.activatePlayer(&p)
    }
}

