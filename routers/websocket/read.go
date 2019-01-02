package websocket

import (
	"time"

	"github.com/haskaalo/intribox/routers/websocket/payload"
)

func (c *Client) readPayloads() {
	// Set Max Message Size
	c.conn.SetReadLimit(512)
	// Close connection if didn't receive pong in period
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// Continue closing connection if message aren't read for too long
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		pl := new(payload.Base)
		err := c.conn.ReadJSON(pl)
		if err != nil {
			if IsJSONError(err) {
				c.close(CloseDecodeError)
			}
			return
		}

		switch pl.Type {

		default:
			c.close(CloseUnknownType)
			return
		}
	}
}
