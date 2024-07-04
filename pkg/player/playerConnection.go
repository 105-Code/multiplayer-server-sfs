package player

import (
	"log"
	"sync"
	"time"

	"github.com/105-Code/multiplayer-server-sfs/pkg/math"
	"github.com/105-Code/multiplayer-server-sfs/pkg/transport"
	"github.com/google/uuid"
)

type PlayerConnection struct {
	Conn          transport.Connection
	MsgFromClient chan []byte
	ExitChan      chan int
	SendMutex     sync.Mutex

	Info PlayerInfo
}

type PlayerInfo struct {
	Id       uuid.UUID
	Position math.Vector2
	Rotation math.Vector2
	Name     string
}

func NewPlayer(playerName string, conn transport.Connection) *PlayerConnection {
	return &PlayerConnection{
		Conn:          conn,
		MsgFromClient: make(chan []byte, 5),
		ExitChan:      make(chan int),
		Info: PlayerInfo{
			Name: playerName,
		},
	}
}

// RecieveMessages ..
func (p *PlayerConnection) RecieveMessages() {
	for {
		_, message, err := p.Conn.ReadMessage()
		p.MsgFromClient <- message
		if err != nil {
			log.Println(p.Info.Name, " has quit.")
			p.Conn.Close()
			// notify that the player has quit
			close(p.ExitChan)
			break
		}
	}
}

// SendMessage ..
func (p *PlayerConnection) SendMessage(msg []byte) error {
	p.SendMutex.Lock()
	defer p.SendMutex.Unlock()

	// log.Println("Sending message to", p.name, ":", string(msg))
	return p.Conn.WriteMessage(1, msg)
}

// Close ..
func (p *PlayerConnection) Close() {
	time.Sleep(1 * time.Second)
	p.Conn.Close()
}
