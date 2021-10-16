package pool

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.handleRegisterClient(client)
			break

		case client := <-pool.Unregister:
			pool.handleUnregisterClient(client)
			break

		case message := <-pool.Broadcast:
			pool.sendMessageToClients(message)
		}
	}
}

func (pool *Pool) sendMessageToClients(message Message) {
	fmt.Println("Sending message to all clients in Pool")

	for client, _ := range pool.Clients {
		if err := client.Conn.WriteJSON(message); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (pool *Pool) handleRegisterClient(client *Client) {
	pool.Clients[client] = true
	fmt.Println("Size of connection pool: ", len(pool.Clients))

	for client, _ := range pool.Clients {
		client.Conn.WriteJSON(Message{
			Type: 1,
			Body: "New user joined ...",
		})
	}
}

func (pool *Pool) handleUnregisterClient(client *Client) {
	delete(pool.Clients, client)
	fmt.Println("Size of connection pool: ", len(pool.Clients))

	for client, _ := range pool.Clients {
		client.Conn.WriteJSON(Message{
			Type: 1,
			Body: "User disconnected ...",
		})
	}
}
