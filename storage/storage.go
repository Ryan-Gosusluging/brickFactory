package storage

import "github.com/Ryan-Gosusluging/brickFactory/brick"

type Storage struct {
	RedBricks   []brick.Brick
	WhiteBricks []brick.Brick
}

func NewStorage() *Storage {
	return &Storage{}
}