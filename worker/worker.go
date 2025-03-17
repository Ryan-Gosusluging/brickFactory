package worker

import (
	"fmt"
	"github.com/Ryan-Gosusluging/brickFactory/factory"
	"github.com/Ryan-Gosusluging/brickFactory/storage"
)

// Worker представляет рабочего, который перемещает кирпичи на склад.
type Worker struct {
	ID           int
	RedStorage   chan factory.Brick
	WhiteStorage chan factory.Brick
}

// NewWorker создает нового рабочего с каналами для складов.
func NewWorker(id int) *Worker {
	return &Worker{
		ID:           id,
		RedStorage:   make(chan factory.Brick, 100),
		WhiteStorage: make(chan factory.Brick, 100),
	}
}

// ReceiveBrick принимает кирпич и отправляет его на соответствующий склад.
func (w *Worker) ReceiveBrick(brick factory.Brick) {
	if brick.Color == "red" {
		w.RedStorage <- brick
		fmt.Printf("Рабочий %d отправил красный кирпич на склад.\n", w.ID)
	} else {
		w.WhiteStorage <- brick
		fmt.Printf("Рабочий %d отправил белый кирпич на склад.\n", w.ID)
	}
}