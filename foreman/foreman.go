package foreman

import (
	"fmt"
	"github.com/Ryan-Gosusluging/brickFactory/factory"
	"github.com/Ryan-Gosusluging/brickFactory/worker"
)

type Foreman struct {
	Workers []*worker.Worker
}

func NewForeman(numWorkers int) *Foreman {
	foreman := &Foreman{}
	for i := 0; i < numWorkers; i++ {
		foreman.Workers = append(foreman.Workers, worker.NewWorker(i+1))
	}
	return foreman
}

func (f *Foreman) DistributeBricks(redChan, whiteChan <-chan factory.Brick) {
	workerIndex := 0
	for {
		select {
		case brick := <-redChan:
			fmt.Printf("Прораб получил красный кирпич.\n")
			f.Workers[workerIndex%len(f.Workers)].ReceiveBrick(brick) 
			workerIndex++
		case brick := <-whiteChan:
			fmt.Printf("Прораб получил белый кирпич.\n")
			f.Workers[workerIndex%len(f.Workers)].ReceiveBrick(brick)
			workerIndex++
		}
	}
}

