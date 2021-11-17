package websocket

type Pool struct {
	Clients    map[*Client]bool
	Register   chan *Client
	UnRegister chan *Client
	Broadcast  chan *Message
}

func NewPool() *Pool {
	return &Pool{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients = make(map[*Client]bool) // remove later
			p.Clients[client] = true
			for c, _ := range p.Clients {
				c.Write(&Message{
					Type: "join",
					Body: "New user joined",
				})
			}
		case client := <-p.UnRegister:
			delete(p.Clients, client)
			for c, _ := range p.Clients {
				c.Write(&Message{
					Type: "leave",
					Body: "New user joined",
				})
			}
		case message := <-p.Broadcast:
			for c, _ := range p.Clients {
				c.Write(message)
			}
		}
	}
}
