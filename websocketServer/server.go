package websocketServer

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/edoardottt/gochanges/db"
	"golang.org/x/net/websocket"
)

type WebsocketServer struct {
	List []*websocket.Conn
}

// Starts the Websocket server, adding handlers.
func StartWebsocketServer() *WebsocketServer {
	serverInstance := WebsocketServer{}
	http.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
		log.Printf("Added a new websocket connection")
		// TODO: Protect this with a mutex
		serverInstance.List = append(serverInstance.List, ws)
		// Spin to keep alive, as per https://stackoverflow.com/questions/23607811/how-to-store-websocket-connection-in-go
		for {
			var s string
			if err := websocket.Message.Receive(ws, &s); err != nil {
				break
			}
		}
	}))
	return &serverInstance
}

func (srv *WebsocketServer) OnScrapeFunction(scrapeTarget db.ScrapeTarget) {
	log.Printf("Sending scraped target to %d Websockets", len(srv.List))
	for _, ws := range srv.List {
		jsonData, err := json.Marshal(scrapeTarget)
		if err != nil {
			log.Printf("Could not serialize scrape target to send to Websocket")
		}
		ws.Write(jsonData)
	}
}
