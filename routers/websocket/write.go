package websocket

import (
	"time"

	ws "github.com/gorilla/websocket"
)

func (c *Client) writePayloads() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		c.close([]byte{})
		ticker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(ws.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
