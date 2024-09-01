package webSocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// メッセージ構造体
type Message struct {
	Message string
	RoomID  string
}

// ルームを取得するか、新規作成
func getOrCreateRoom(roomID string) *Room {
	room, exists := rooms[roomID]
	if !exists {
		room = &Room{
			clients:   make(map[*websocket.Conn]bool),
			broadcast: make(chan Message),
		}
		rooms[roomID] = room

		// ブロードキャスト処理をゴルーチンで開始
		go room.handleMessages()
	}
	return room
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {

		}
	}(ws)

	// クエリパラメータからルームIDを取得
	roomID := r.URL.Query().Get("room_id")
	if roomID == "" {
		roomID = "default"
	}

	// ルームを取得または作成
	room := getOrCreateRoom(roomID)

	room.mu.Lock()
	room.clients[ws] = true
	room.mu.Unlock()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("メッセージの受信エラー: %v", err)
			room.mu.Lock()
			delete(room.clients, ws)
			room.mu.Unlock()
			break
		}
		// 受信したメッセージをブロードキャストチャネルに送信
		room.broadcast <- Message{string(message), roomID}
	}
}

func (room *Room) handleMessages() {
	for {
		msg := <-room.broadcast
		room.mu.Lock()
		for client := range room.clients {
			go func(client *websocket.Conn, msg Message) { // 非同期でメッセージを送信
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("メッセージの送信エラー: %v", err)
					err := client.Close()
					if err != nil {
						return
					}
					delete(room.clients, client)
				}
			}(client, msg)
		}
		room.mu.Unlock()
	}
}
