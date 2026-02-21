package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"log/slog"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Packet struct {
	Value1 int32  `json:"value1"`
	Value2 uint32 `json:"value2"`
}

// ===== WebSocket 管理 =====

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex

func broadcast(p Packet) {
	mu.Lock()
	defer mu.Unlock()

	data, _ := json.Marshal(p)

	for c := range clients {
		err := c.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			c.Close()
			delete(clients, c)
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	slog.Info("New WebSocket client connected", "clients", len(clients), "remoteAddr", conn.RemoteAddr().String())
}

// ===== UDP受信 =====

func udpReceiver() {
	addr, err := net.ResolveUDPAddr("udp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		buf := make([]byte, 324)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil || n < 8 {
			continue
		}

		var p Packet
		err = binary.Read(
			bytes.NewReader(buf),
			binary.LittleEndian,
			&p,
		)
		if err != nil {
			continue
		}

		broadcast(p)
	}
}

func main() {
	go udpReceiver()

	http.HandleFunc("/ws", wsHandler)

	log.Println("Server running :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
