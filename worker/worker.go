package worker

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"github.com/Ryan-Gosusluging/brickFactory/factory"
)

func ManageWorker(id int, dispatch <-chan factory.Brick, redStorage chan<- factory.Brick, whiteStorage chan<- factory.Brick, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Рабочий %d: Приступил к работе\n", id)
	defer fmt.Printf("Рабочий %d: Завершил работу\n", id)

	for brick := range dispatch {
		fmt.Printf("Рабочий %d: получил %s кирпич\n", id, brick.Color)
		if brick.Color == "red" {
			redStorage <- brick
			fmt.Printf("Рабочий %d: Отправил красный кирпич на склад\n", id)
		} else {
			whiteStorage <- brick
			fmt.Printf("Рабочий %d: Отправил белый кирпич на склад\n", id)
		}
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}