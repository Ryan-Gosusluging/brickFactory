package factory

import (
	"fmt"
	"math/rand"
	"time"
)

type Brick struct {
	Color string
}

func ProduceBricks(redChan, whiteChan chan<- Brick) {
	fmt.Println("Фабрика: Производство запущено!")
	rand.Seed(time.Now().UnixNano()) //цвет кирпича выбирается случайным образом
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
		time.Sleep(time.Second)
	}
}