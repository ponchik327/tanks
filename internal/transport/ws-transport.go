package transport

import (
	"github.com/gorilla/websocket"
)

// WebSocketTransport реализует Transport для WebSocket
type WebSocketTransport struct {
	conn *websocket.Conn
}

func NewWebSocketTransport() *WebSocketTransport {
	return &WebSocketTransport{}
}

func (t *WebSocketTransport) Connect(address string) error {
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(address, nil)
	if err != nil {
		return err
	}
	t.conn = conn
	return nil
}

func (t *WebSocketTransport) Send(data []byte) error {
	return t.conn.WriteMessage(websocket.BinaryMessage, data)
}

func (t *WebSocketTransport) Receive() ([]byte, error) {
	_, message, err := t.conn.ReadMessage()
	return message, err
}

func (t *WebSocketTransport) Close() error {
	return t.conn.Close()
}

func (t *WebSocketTransport) IsConnected() bool {
	return t.conn != nil
}
