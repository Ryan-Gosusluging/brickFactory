package worker

import (
	"fmt"
	"github.com/Ryan-Gosusluging/brickFactory/factory"
)

type Worker struct {
	ID           int
	RedStorage   chan factory.Brick
	WhiteStorage chan factory.Brick
}

func NewWorker(id int) *Worker {
	return &Worker{
		ID:           id,
		RedStorage:   make(chan factory.Brick, 100),
		WhiteStorage: make(chan factory.Brick, 100),
	}
}

func (w *Worker) ReceiveBrick(brick factory.Brick) {
	fmt.Printf("Рабочий %d: Приступил к работе\n", w.ID)
	if brick.Color == "red" {
		w.RedStorage <- brick
		fmt.Printf("Рабочий %d отправил красный кирпич на склад.\n", w.ID)
	} else {
		w.WhiteStorage <- brick
		fmt.Printf("Рабочий %d отправил белый кирпич на склад.\n", w.ID)
	}
}