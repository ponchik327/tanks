package core

import (
	"tank.io/internal/logic"
	"tank.io/internal/transport"
)

type GameServer struct {
	transport transport.Transport
}

func NewGameServer(transport transport.Transport) *GameServer {
	return &GameServer{
		transport: transport,
	}
}

func (s *GameServer) Start(address string) error {
	if err := s.transport.Connect(address); err != nil {
		return err
	}

	defer s.transport.Close()

	// Пример игрового цикла
	for s.transport.IsConnected() {
		data, err := s.transport.Receive()
		if err != nil {
			return err
		}

		// Обработка данных
		response := logic.ProcessGameData(data)

		// Отправка ответа
		if err := s.transport.Send(response); err != nil {
			return err
		}
	}

	return nil
}
