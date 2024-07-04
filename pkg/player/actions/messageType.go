package actions

type MessageType int

const (
	Ping MessageType = iota + 1
	PlayerUpdate
	EnterWorld
	LeaveWorld
	UpdateWorld // just to send to players
)
