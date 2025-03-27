package factory

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/Ryan-Gosusluging/brickFactory/brick"
)

type Factory struct {
    RedChan   chan<- brick.Brick
    WhiteChan chan<- brick.Brick
}

func NewFactory(redChan, whiteChan chan<- brick.Brick) *Factory {
    return &Factory{
        RedChan:   redChan,
        WhiteChan: whiteChan,
    }
}

func (f *Factory) ProduceBricks(done <-chan struct{}) {
	rand.Seed(time.Now().UnixNano())
    for {
        select {
        case <-done:
            return
        default:
            color := "red"
            if rand.Intn(2) == 1 {
                color = "white"
            }
            
            brick := brick.Brick{Color: color}
            if color == "red" {
                f.RedChan <- brick
            } else {
                f.WhiteChan <- brick
            }
            time.Sleep(500 * time.Millisecond)
        }
    }
}