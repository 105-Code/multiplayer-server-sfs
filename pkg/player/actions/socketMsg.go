package actions

import (
	"encoding/json"
	"log"
)

type SocketMsg struct {
	Type    MessageType     `json:"type"`
	Message json.RawMessage `json:"msg"`
}

func ParseSocketMsg(msgBytes []byte) *SocketMsg {
	msg := SocketMsg{}
	err := json.Unmarshal(msgBytes, &msg)
	if err != nil {
		log.Println("ERROR", "Error in recieving message", err)
	}
	return &msg
}

// ToBytes returns the socket message in bytes
func (msg *SocketMsg) ToBytes() (returnMsg []byte) {
	returnMsg, err := json.Marshal(msg)
	if err != nil {
		returnMsg = []byte(err.Error())
	}
	return
}
