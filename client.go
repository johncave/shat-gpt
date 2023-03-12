package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// IncomingPress is the type used to read the JSON coming down the websocket from the client
type IncomingPress struct {
	EventName string `json:"event_name"`
	UserToken string `json:"token"`
}

type OutgoingUpdate struct {
	EventName string    `json:"event_name"`
	PoolSize  int64     `json:"pool_size"`
	Timestamp time.Time `json:"timestamp"`
}

// readPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump() {
	c := s.conn
	defer func() {
		h.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var incomingJson IncomingPress
		err := c.ws.ReadJSON(incomingJson)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		// Check what type of message we are receiving
		switch incomingJson.EventName {
		case "press":
			log.Println("Processing incoming press")
			processPress(incomingJson, s.room)
		}

		// m := message{msg, s.room}
		// h.broadcast <- m
	}
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request, roomId string) {
	fmt.Print(roomId)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	s := subscription{c, roomId}
	h.register <- s
	go s.writePump()
	go s.readPump()
}

func processPress(in IncomingPress, room string) {
	// Get the user from the token
	u, err := LookupToken(in.UserToken)
	if err != nil {
		log.Println("User token was invalid - ignoring websocket message", in)
		return
	}

	// User is okay now
	_, err = redisIncrement("global-total")
	if err != nil {
		log.Println("Error incrementing the global total of clicks - skipping", err)
		return
	}
	poolCount, err := redisIncrement("pool-total")
	if err != nil {
		log.Println("Error updating the pool size - skipping", err)
		return
	}

	// Add the press to the leaderboard
	lbUser, err := GlobalLeaderBoard.GetMember(u.UserName)
	logError("updating leaderboard for user", err)
	GlobalLeaderBoard.RankMember(u.UserName, lbUser.Score+1)

	out := OutgoingUpdate{
		EventName: "pool_size",
		PoolSize:  poolCount,
		Timestamp: time.Now(),
	}

	outjson, _ := json.MarshalIndent(out, "", "  ")

	m := message{outjson, room}
	h.broadcast <- m
}
