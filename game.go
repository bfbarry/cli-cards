package main

import (
	"fmt"
	"time"
	"net"
	"sync"
	"math/rand"
	"bufio"
)

type GameState struct {
    players map[Pid]Player
    state int
    deck []Card
	pile []Card
	mu sync.Mutex
	// gameStarted bool
}

func (self *GameState) waitForPlayers(conn net.Conn) {
	var err error
	for len(self.players) < 2 {
		_, err = fmt.Fprintf(conn, "%s", "Waiting for more players...")
		_, err = fmt.Fprintf(conn, "\r")
		time.Sleep(100*time.Millisecond)
	}
	fmt.Fprintf(conn, "Let's get this party started!\n")
	if err != nil {
        fmt.Println("printf err", err)
    }
}

func (self *GameState) createPlayer(ln net.Listener) (Player, error) {
	var p Player
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Connection error:", err)
		return p, err
	}
	init_hand := []Card{}
	id := Pid(rand.Intn(1000))
	p = Player{id: id, hand: init_hand, conn: conn}
	
	self.mu.Lock()
	p.deal(self, 5)
	self.players[id] = p
	self.mu.Unlock()
	
	return p, nil
}

func (self *GameState) rmPlayer(id Pid) {
	delete(self.players, id)
}

func (self *GameState) activatePlayer(player *Player) {
    var err error
    var wg sync.WaitGroup

    defer player.conn.Close()
    fmt.Println("Client connected:", player.conn.RemoteAddr())
    
    self.waitForPlayers(player.conn)
	fmt.Fprintf(player.conn, "> ")
    scanner := bufio.NewScanner(player.conn)
    wg.Add(1)
    go func() {
        defer wg.Done()
        for scanner.Scan() {
            // self.waitForPlayers(player.conn) // TODO should move this outside somewhere
            
            text := scanner.Text()
			
            if text == "show" {
				player.showCards()
			} else if text == "bs" {
				for _, p := range(self.players) {
					fmt.Fprintf(p.conn, "Player %d has called bs!\n", player.id)
				}
				// check if bs condition is true, then last player gets all cards
			} else if len(text) > 0 && text[:len("put")] == "put" {
				// cardString := text[4:]
			}
			fmt.Fprintf(player.conn, "> ")
        }
    }()
    wg.Wait()

    if err = scanner.Err(); err != nil {
        fmt.Println("Connection error:", err)
    }

    self.rmPlayer(player.id)
    fmt.Println("Client disconnected:", player.conn.RemoteAddr())

}