package player

import (
	"fmt"
	"testing"

	mockTransport "github.com/105-Code/multiplayer-server-sfs/pkg/transport/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPlayerConnection_RecieveMessages(t *testing.T) {
	assertion := assert.New(t)
	tests := []struct {
		name            string
		player          string
		expectedMessage []byte
		connSetup       func(c *mockTransport.MockConnection)
	}{
		{
			name:            "Happy path",
			player:          "test 1",
			expectedMessage: []byte("test"),
			connSetup: func(c *mockTransport.MockConnection) {
				c.EXPECT().ReadMessage().Return(1, []byte("test"), nil)
				c.EXPECT().ReadMessage().Return(1, nil, fmt.Errorf("end"))
				c.EXPECT().Close().Return(nil)
			},
		},
		{
			name:   "Fail read message",
			player: "test 1",

			connSetup: func(c *mockTransport.MockConnection) {
				c.EXPECT().ReadMessage().Return(1, nil, fmt.Errorf("error"))
				c.EXPECT().Close().Return(nil)
			},
		},
	}

	ctrl := gomock.NewController(t)
	conn := mockTransport.NewMockConnection(ctrl)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if test.connSetup != nil {
				test.connSetup(conn)
			}

			player := NewPlayer(test.player, conn)

			player.RecieveMessages()

			assertion.Equal(test.expectedMessage, <-player.MsgFromClient)

		})
	}

}

func TestPlayerConnection_SendMessage(t *testing.T) {
	assertion := assert.New(t)
	tests := []struct {
		name        string
		player      string
		message     []byte
		expectError bool
		connSetup   func(c *mockTransport.MockConnection)
	}{
		{
			name:    "Happy path",
			player:  "test 1",
			message: []byte("test"),
			connSetup: func(c *mockTransport.MockConnection) {
				c.EXPECT().WriteMessage(1, []byte("test")).Return(nil)
			},
		},
		{
			name:        "Fail send message",
			player:      "test 1",
			message:     []byte("test"),
			expectError: true,
			connSetup: func(c *mockTransport.MockConnection) {
				c.EXPECT().WriteMessage(1, []byte("test")).Return(fmt.Errorf("fail send message"))
			},
		},
	}

	ctrl := gomock.NewController(t)
	conn := mockTransport.NewMockConnection(ctrl)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if test.connSetup != nil {
				test.connSetup(conn)
			}

			player := NewPlayer(test.player, conn)

			err := player.SendMessage(test.message)

			if test.expectError {
				assertion.NotNil(err)
			} else {
				assertion.Nil(err)
			}

		})
	}

}
