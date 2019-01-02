package websocket

import (
	"time"

	ws "github.com/gorilla/websocket"
	"github.com/haskaalo/intribox/models"
)

// Client socket
type Client struct {
	// The websocket connection
	conn *ws.Conn

	// The user session info
	session *models.Session

	// Sent when client connection died/ended
	end chan bool
}

func (c *Client) close(message []byte) {
	c.conn.WriteControl(ws.CloseMessage, message, time.Now().Add(writeWait))
	c.conn.Close()
	c.end <- true
}

func (c *Client) sendJSON(message interface{}) error {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	err := c.conn.WriteJSON(message)
	return err
}
