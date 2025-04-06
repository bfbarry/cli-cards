package main

import (
    "bufio"
    "fmt"
    "net"
    "math/rand"
)


type Player struct {
    id string
    hand []Card    
}

type GameState struct {
    players []Player
    state int
    deck []Card
}

func handleConnection(conn net.Conn, state *GameState) {
    defer conn.Close()
    fmt.Println("Client connected:", conn.RemoteAddr())
    
    init_hand := []Card{}
    p := Player{id: "stuff", hand: init_hand}
    p.deal(state, 5)
    state.players = append(state.players, p)
    
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        text := scanner.Text()
        fmt.Printf("Received: %s\n", text)
        p.showCards(conn)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Connection error:", err)
    }

    fmt.Println("Client disconnected:", conn.RemoteAddr())
}

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting TCP server:", err)
        return
    }
    defer ln.Close()

    fmt.Println("TCP server listening on :8080")

    init_players := []Player{}
    deck := init_deck()
    rand.Shuffle(len(deck), func(i, j int) {
        deck[i], deck[j] = deck[j], deck[i]
    })

    game := GameState{players: init_players, state: 0, deck: deck}

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Connection error:", err)
            continue
        }

        go handleConnection(conn, &game)
    }
}

