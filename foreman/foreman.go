package foreman

import (
	"fmt"
	"time"
	"github.com/Ryan-Gosusluging/brickFactory/factory"
)

func DistributionBricks(bricksIn <-chan factory.Brick, dispatch chan<- factory.Brick, dispatchRate int){
	fmt.Println("Прораб: Начало работы...")
	defer fmt.Println("Прораб: Работа завершена.")
	for brick := range bricksIn {
		fmt.Printf("Прораб: Получил %s кирпич и отправляю рабочим\n", brick.Color)
		dispatch <- brick
		time.Sleep(time.Second / time.Duration(dispatchRate))
	}
	close(dispatch)
}