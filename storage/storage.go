package storage

import "github.com/Ryan-Gosusluging/brickFactory/factory"

type Storage struct {
	RedStorage   chan factory.Brick
	WhiteStorage chan factory.Brick
}

func NewStorage() *Storage {
	return &Storage{
		RedStorage:   make(chan factory.Brick, 100),
		WhiteStorage: make(chan factory.Brick, 100),
	}
}