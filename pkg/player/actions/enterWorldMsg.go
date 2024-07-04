package actions

import "github.com/google/uuid"

// PlayerUpdate is used by both client and server to notify about player change
type WnterWorldMsg struct {
	id uuid.UUID
}
