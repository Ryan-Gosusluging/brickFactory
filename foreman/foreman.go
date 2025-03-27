package foreman

import (
	"fmt"
	"github.com/Ryan-Gosusluging/brickFactory/brick"
	"github.com/Ryan-Gosusluging/brickFactory/worker"
	"sync"
)

type Foreman struct {
	Workers []*worker.Worker
	BrickChan chan brick.Brick
    mu        sync.Mutex
    bricks    []brick.Brick
    done      chan struct{}
}

func NewForeman(numWorkers int, redStorage, whiteStorage chan<- brick.Brick) *Foreman {
    f := &Foreman{
        BrickChan: make(chan brick.Brick, 100),
        done:      make(chan struct{}),
    }
    for i := 0; i < numWorkers; i++ {
        f.Workers = append(f.Workers, worker.NewWorker(i+1, redStorage, whiteStorage))
    }
    return f
}

func (f *Foreman) DistributeBricks(wg *sync.WaitGroup) {
	defer wg.Done()
    
    // Собираем все кирпичи
    for brick := range f.BrickChan {
        f.mu.Lock()
        f.bricks = append(f.bricks, brick)
        
        // Проверка на 70% белых кирпичей
        if f.whitePercentage() > 70 {
            f.mu.Unlock()
            close(f.done)
            return
        }
        f.mu.Unlock()
    }
    
    // Распределяем кирпичи по рабочим
    for i, brick := range f.bricks {
        workerIdx := i % len(f.Workers)
        f.Workers[workerIdx].BrickChan <- brick
    }
    
    // Закрываем каналы рабочих
    for _, w := range f.Workers {
        close(w.BrickChan)
    }
}

func (f *Foreman) whitePercentage() float64 {
    white := 0
    for _, b := range f.bricks {
        if b.Color == "white" {
            white++
        }
    }
    if len(f.bricks) == 0 {
        return 0
    }
    return float64(white) / float64(len(f.bricks)) * 100
}

