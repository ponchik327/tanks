package transport

type Transport interface {
	// Connect устанавливает соединение
	Connect(address string) error

	// Send отправляет данные
	Send(data []byte) error

	// Receive получает данные
	Receive() ([]byte, error)

	// Close закрывает соединение
	Close() error

	// IsConnected проверяет состояние соединения
	IsConnected() bool
}
