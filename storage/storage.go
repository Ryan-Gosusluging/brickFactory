package storage

import "github.com/Ryan-Gosusluging/brickFactory/factory"

// Storage представляет склады для красных и белых кирпичей.
type Storage struct {
	RedStorage   chan factory.Brick
	WhiteStorage chan factory.Brick
}

// NewStorage создает новый склад.
func NewStorage() *Storage {
	return &Storage{
		RedStorage:   make(chan factory.Brick, 100),
		WhiteStorage: make(chan factory.Brick, 100),
	}
}