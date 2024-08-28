package usecase

// WebSocketUseCase はWebSocketのユースケースを表します。
type WebSocketUseCase interface {
	Open(message []byte) ([]byte, error)
}

type WebSocket struct {
}

var _ WebSocketUseCase = (*WebSocket)(nil)

// NewWebSocketUseCase はWebSocketUseCaseを生成します。
func NewWebSocketUseCase() *WebSocket {
	return &WebSocket{}
}

func (w *WebSocket) Open(message []byte) ([]byte, error) {
	// ここにWebSocketを開く処理を書く
	message = append(message, []byte("**************")...)
	return message, nil
}
