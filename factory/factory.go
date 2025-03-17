package factory

import (
	"math/rand"
	"time"
)

// Brick представляет кирпич с определенным цветом.
type Brick struct {
	Color string
}

// ProduceBricks производит кирпичи и отправляет их в соответствующие каналы.
func ProduceBricks(redChan, whiteChan chan<- Brick) {
	rand.Seed(time.Now().UnixNano())
	for {
		color := "red"
		if rand.Intn(2) == 1 {
			color = "white"
		}
		brick := Brick{Color: color}
		if color == "red" {
			redChan <- brick
			fmt.Printf("Произведен красный кирпич.\n")
		} else {
			whiteChan <- brick
			fmt.Printf("Произведен белый кирпич.\n")
		}
		time.Sleep(time.Second) // Имитация времени производства
	}
}