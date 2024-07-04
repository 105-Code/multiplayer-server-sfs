// spurce code: https://github.com/aarthikrao/jumpAndShoot
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/105-Code/multiplayer-server-sfs/pkg/config"
	"github.com/105-Code/multiplayer-server-sfs/pkg/gameserver"
	"github.com/105-Code/multiplayer-server-sfs/pkg/logger"
	"github.com/105-Code/multiplayer-server-sfs/pkg/player"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var Server *gameserver.GameServer

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

func main() {
	Server = gameserver.GetGameServer(time.Duration(config.AppConfig.Port) * time.Millisecond)

	http.HandleFunc("/game", handler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "JumpAndShoot engine is running on this port")
	})

	var address = fmt.Sprintf("127.0.0.1:%d", config.AppConfig.Port)

	logger.Info("Start listening at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		logger.Error("Fail starting server %s", err)
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	playerName := r.URL.Query()["id"][0]
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ERROR", err)
		return
	}
	logger.Info("Join", playerName)
	// Create a new player instance
	p := player.NewPlayer(playerName, conn)
	p.Info.Id = uuid.New()
	Server.AddClient(p)
}
