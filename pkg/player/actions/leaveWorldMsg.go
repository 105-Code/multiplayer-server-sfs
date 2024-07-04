package actions

import "github.com/google/uuid"

// PlayerUpdate is used by both client and server to notify about player change
type LeaveWorldMsg struct {
	id uuid.UUID
}
