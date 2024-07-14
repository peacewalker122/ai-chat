package chat

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/gorilla/websocket"
	"google.golang.org/api/iterator"
)

type Handler struct {
	*genai.GenerativeModel
}

func NewHandler(mdl *genai.GenerativeModel) *Handler {
	return &Handler{mdl}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *Handler) Chat(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()

	// Set up a ping-pong mechanism with timeout
	pingPeriod := 5 * time.Second
	pongWait := 10 * time.Second

	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	go func() {
		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					log.Println("write ping:", err)
					conn.SetReadDeadline(time.Now().Add(pongWait))
					return
				}
			}
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		resp := h.GenerateContentStream(context.Background(), genai.Text(string(message)))
		for {
			msg, err := resp.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				break
			}

			conn.SetWriteDeadline(time.Now().Add(pongWait))
			for _, candidate := range msg.Candidates {
				if candidate.Content != nil {
					content := candidate.Content

					data, err := json.Marshal(content)
					if err != nil {
						conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
						return
					}

					if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
						log.Println("write:", err)
						return
					}
				}
			}
		}
	}
}
