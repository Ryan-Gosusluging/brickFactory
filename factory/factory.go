package factory

import (
	"fmt"
	"math/rand"
	"time"
)

type Brick struct{
	Color string
}

func ProduceBrick(bricksOut chan<- Brick, productionRate int, simulationTimeSec int){
	fmt.Println("Фабрика: Производство запущено!")
	defer fmt.Println("Фабрика: Производство завершено!")
	startTime := time.Now()
	for time.Since(startTime).Seconds() < float64(simulationTimeSec){
		color := "red"
		if rand.Intn(2) == 0{
			color = "white"
		}
		brick := Brick{Color: color}
		bricksOut <- brick
		fmt.Println("Фабрика: произведен %s кирпич\n", color)
		time.Sleep(time.Seconds / time.Duration(productionRate))
	}
	close(bricksOut)
}