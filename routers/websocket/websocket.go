package websocket

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	ws "github.com/gorilla/websocket"
	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/routers/websocket/payload"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// ServeWs Handler
func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn().AnErr("error", err).Msg("Failed to upgrade a connection")
		return
	}

	// Read initial message with a deadline
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	pl := new(payload.Base)
	err = conn.ReadJSON(pl)
	if err != nil {
		conn.WriteControl(ws.CloseMessage, CloseAuthenticationFailed, time.Now().Add(writeWait))
		conn.Close()
		return
	}

	data := new(payload.ReceiveAuth)
	err = json.Unmarshal(pl.Data, data)
	if err != nil {
		conn.WriteControl(ws.CloseMessage, CloseDecodeError, time.Now().Add(writeWait))
		conn.Close()
		return
	}

	// Check if token is valid
	session, err := models.GetSessionByToken(data.Token)
	if err == models.ErrNotValidSessionToken {
		conn.WriteControl(ws.CloseMessage, CloseAuthenticationFailed, time.Now().Add(writeWait))
		conn.Close()
		return
	} else if err != nil {
		log.Error().AnErr("error", err).Msg("Failed to GetSessionByToken")
		conn.WriteControl(ws.CloseMessage, CloseUnknownError, time.Now().Add(writeWait))
		conn.Close()
		return
	}

	client := &Client{conn: conn, session: session, end: make(chan bool)}
	client.sendJSON(&payload.SendAuthSuccess{
		Ok: 1,
	})

	go client.writePayloads()
	go client.readPayloads()
}
