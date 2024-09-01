package webSocket

import (
	"awesomeProject1/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type handler struct {
	u  usecase.WebSocketUseCase
	c  usecase.ChatMessageU
	se *sessions.CookieStore
}

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

// ルームを管理する構造体
type Room struct {
	clients   map[*websocket.Conn]bool // ルームに参加しているクライアント
	broadcast chan Message             // ルーム内でブロードキャストするメッセージ
	mu        sync.Mutex               // 排他制御のためのMutex
}

var rooms = make(map[string]*Room)

func NewRouter(u usecase.WebSocketUseCase, chu usecase.ChatMessageU, se *sessions.CookieStore) http.Handler {

	r := chi.NewRouter()

	//r.Use(aunth.Middleware(se))

	h := &handler{
		u:  u,
		c:  chu,
		se: se,
	}
	r.Get("/chat/open", h.Open)
	r.Get("/chatRoom", handleConnections)
	r.Get("/roomInf", roomCount)

	return r

}
