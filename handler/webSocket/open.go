package webSocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	HandshakeTimeout: 5 * time.Second,               // ハンドシェイクのタイムアウトを5秒に設定
	ReadBufferSize:   1024,                          // 読み取りバッファを1KBに設定
	WriteBufferSize:  1024,                          // 書き込みバッファを1KBに設定
	Subprotocols:     []string{"chat", "superchat"}, // サポートするサブプロトコルを設定
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		http.Error(w, reason.Error(), status)
	},
	EnableCompression: true, // メッセージの圧縮を有効にする
}

func (h *handler) Open(w http.ResponseWriter, r *http.Request) {
	// HTTP接続をWebSocketにアップグレード
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	ctx := r.Context()

	go func() {
		for {
			select {
			case <-ticker.C:
				err := conn.WriteMessage(websocket.TextMessage, []byte("ping"))
				if err != nil {
					log.Println("Write error:", err)
					return
				}
			}
		}
	}()

	for {
		// クライアントからメッセージを読み込み
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received: %s", message)

		if err = h.c.Save(ctx, 1, 1, string(message)); err != nil {
			log.Println("Save error:", err)
		}

		// クライアントにメッセージを送り返す
		message, _ = h.u.Open(message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
