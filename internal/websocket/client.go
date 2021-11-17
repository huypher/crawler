package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/v2/util/log"
)

type Client struct {
	Conn *websocket.Conn
	Pool *Pool
}

func (c *Client) Write(msg *Message) error {
	err := c.Conn.WriteJSON(msg)
	if err != nil {
		log.Errorf("Send message=%v", msg)
		return err
	}

	return nil
}
