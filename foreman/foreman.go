package foreman

import (
	"fmt"
	"github.com/Ryan-Gosusluging/brickFactory/factory"
	"github.com/Ryan-Gosusluging/brickFactory/worker"
)

// Foreman представляет прораба, который распределяет кирпичи между рабочими.
type Foreman struct {
	Workers []*worker.Worker
}

// NewForeman создает нового прораба с указанным количеством рабочих.
func NewForeman(numWorkers int) *Foreman {
	foreman := &Foreman{}
	for i := 0; i < numWorkers; i++ {
		foreman.Workers = append(foreman.Workers, worker.NewWorker(i+1))
	}
	return foreman
}

// DistributeBricks распределяет кирпичи между рабочими.
func (f *Foreman) DistributeBricks(redChan, whiteChan <-chan factory.Brick) {
	for {
		select {
		case brick := <-redChan:
			fmt.Printf("Прораб получил красный кирпич.\n")
			f.Workers[0].ReceiveBrick(brick)
		case brick := <-whiteChan:
			fmt.Printf("Прораб получил белый кирпич.\n")
			f.Workers[1].ReceiveBrick(brick)
		}
	}
}