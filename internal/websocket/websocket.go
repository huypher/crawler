package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Websocket struct {
	Upgrader    websocket.Upgrader
	Pool        *Pool
	httpHandler http.Handler
	Server      *http.Server
}

type Message struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

func NewWebsocketService() *Websocket {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	pool := NewPool()

	server := &http.Server{
		Addr: ":8080",
	}

	return &Websocket{
		Upgrader: upgrader,
		Pool:     pool,
		Server:   server,
	}
}

func (ws *Websocket) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return conn, err
	}
	return conn, nil
}

func (ws *Websocket) ServeWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := ws.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &Client{
		Conn: conn,
		Pool: ws.Pool,
	}

	ws.Pool.Register <- client
}

func (ws *Websocket) AddHandler(handler http.Handler) {
	ws.httpHandler = handler
}

func (ws *Websocket) Start() {
	go ws.Pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(w, r)
	})
	ws.Server.ListenAndServe()
}

func (ws *Websocket) Broadcast(message *Message) {
	ws.Pool.Broadcast <- message
}
