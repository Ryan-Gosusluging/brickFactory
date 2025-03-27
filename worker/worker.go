package worker

import (
	"fmt"
	"github.com/Ryan-Gosusluging/brickFactory/brick"
	"sync"
)

type Worker struct {
	ID int
	BrickChan chan brick.Brick
	RedStorage   chan<- brick.Brick
	WhiteStorage chan<- brick.Brick
}

func NewWorker(id int, redStorage, whiteStorage chan<- brick.Brick) *Worker {
	return &Worker{
        ID: id,
        BrickChan:  make(chan brick.Brick, 10),
        RedStorage:   redStorage,
        WhiteStorage: whiteStorage,
    }
}

func (w *Worker) ReceiveBrick(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Рабочий %d: Приступил к работе\n", w.ID)
	for brick := range w.BrickChan{
		if brick.Color == "red" {
			w.RedStorage <- brick
			fmt.Printf("Рабочий %d отправил красный кирпич на склад.\n", w.ID)
		} else {
			w.WhiteStorage <- brick
			fmt.Printf("Рабочий %d отправил белый кирпич на склад.\n", w.ID)
		}
	}
}