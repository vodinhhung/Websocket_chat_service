package main

import (
	"fmt"
	"net/http"

	"github.com/hungvo/chatservice/backend/common/pool"
	"github.com/hungvo/chatservice/backend/common/websocketutil"
)

func setupRoutes() {
	pool := pool.NewPool()
	go pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple server")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func serveWs(pooling *pool.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")

	conn, err := websocketutil.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &pool.Client{
		Conn: conn,
		Pool: pooling,
	}

	pooling.Register <- client
	client.Read()
}

func main() {
	fmt.Println("Start server")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
