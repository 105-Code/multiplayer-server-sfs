package gameserver

import (
	"sync"
	"time"

	"github.com/105-Code/multiplayer-server-sfs/pkg/config"
	"github.com/105-Code/multiplayer-server-sfs/pkg/logger"
	"github.com/105-Code/multiplayer-server-sfs/pkg/player"
	"github.com/105-Code/multiplayer-server-sfs/pkg/player/actions"
	"github.com/105-Code/multiplayer-server-sfs/pkg/rocket"
)

type GameServer struct {
	clientMap   map[string]*player.PlayerConnection
	debris      []rocket.Rocket // uncontroller objects. ej first stages, abandonded rockets.
	connections int
	tick_rate   time.Duration
	mu          sync.Mutex
}

func GetGameServer(tick_rate time.Duration) *GameServer {
	return &GameServer{
		clientMap: make(map[string]*player.PlayerConnection),
		tick_rate: tick_rate,
		debris:    []rocket.Rocket{},
	}
}

func (gs *GameServer) AddClient(p *player.PlayerConnection) {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	if gs.connections+1 > config.AppConfig.TotalPlayers {
		return
	}

	gs.clientMap[p.Info.Name] = p
	// Start listening to messages from player
	gs.connections += 1
	go p.RecieveMessages()
	go gs.routeMessage(p)
}

func (gs *GameServer) RemoveClient(p *player.PlayerConnection) {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	gs.connections -= 1
	delete(gs.clientMap, p.Info.Name)
}

func (gs *GameServer) Brodcast(msg *actions.SocketMsg) {
	var message = msg.ToBytes()
	gs.mu.Lock()
	defer gs.mu.Unlock()
	for _, player := range gs.clientMap {
		player.SendMessage(message)
	}
}

func (gs *GameServer) routeMessage(player *player.PlayerConnection) {
BREAK:
	for {
		select {
		case msgBytes := <-player.MsgFromClient:
			msg := actions.ParseSocketMsg(msgBytes)
			gs.handleActionMessage(msg, player)

		// Handle exit messages
		case <-player.ExitChan:
			logger.Info("%s has exited the game", player.Info.Name)
			gs.RemoveClient(player)
			go player.Close()
			break BREAK
		}
	}
}

func (gs *GameServer) handleActionMessage(msg *actions.SocketMsg, actor *player.PlayerConnection) {
	switch msg.Type {
	case actions.Ping:
		// Ping the message back to client
		actor.SendMessage(msg.ToBytes())
	case actions.PlayerUpdate:
		//handlePositionUpdate(msg, player, opponent)

	}
}

/*
func handlePositionUpdate(msg *models.SocketMessage, player *Player, opponent *Player) {
	var pU models.PlayerUpdate
	err := json.Unmarshal(msg.Message, &pU)
	if err != nil {
		log.Println("ERROR", "Invalid message", msg.Message, err)
	}
	if pU.Fire && !player.ready {
		// The player is ready for the match
		log.Println(player.name, "is ready for the match")
		player.ready = true
	}

	// Update the player position value in server
	player.posY = pU.PlayerPositionY

	// Marshal the json and send to player
	j, _ := json.Marshal(pU)
	msg.Message = j
	player.SendMessage(msg.ToBytes())

	// Marshal the json and send to opponent
	pU.IsOpponent = true
	j, _ = json.Marshal(pU)
	msg.Message = j
	opponent.SendMessage(msg.ToBytes())
}

func handleCollisionMessage(msg *models.SocketMessage, player *Player, opponent *Player) {
	if !player.ready || !opponent.ready {
		return
	}
	// verify message
	var cM models.CollisionRequest
	err := json.Unmarshal(msg.Message, &cM)
	if err != nil {
		log.Println("ERROR", "Error in parsing collision message", err)
	}
	if cM.Character == 1 {
		// Player has collided
		log.Println("Collision detected: " + player.name)
		player.lives--
		psu := models.GetScoreUpdateSocketBytes(player.lives, opponent.lives)
		player.SendMessage(psu)

		osu := models.GetScoreUpdateSocketBytes(opponent.lives, player.lives)
		opponent.SendMessage(osu)
	}
	if player.lives <= 0 {
		msg := getGameEndMessage(opponent.name)
		player.SendMessage(msg)
		opponent.SendMessage(msg)

		go player.Close()
		go opponent.Close()
	}
}

func getGameEndMessage(winner string) []byte {
	gW := models.GameEnd{
		Winner: winner,
	}
	winMessage, _ := json.Marshal(gW)
	msg := models.SocketMessage{
		Type:    models.GameEndMsg,
		Message: winMessage,
	}
	return msg.ToBytes()
}
*/
