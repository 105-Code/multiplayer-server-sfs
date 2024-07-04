package actions

import "github.com/105-Code/multiplayer-server-sfs/pkg/math"

// PlayerUpdate is used by both client and server to notify about player change
type PlayerUpdateMsg struct {
	Position math.Vector2 `json:"position"`
	Rotation math.Vector2 `json:"rotation"`
}
