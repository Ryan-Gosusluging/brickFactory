package storage

import (
	"fmt"
	"sync"
	"time"

	"example.com/brickfactory/factory"
)

func ManageStorage(name string, storageIn <-chan factory.Brick, cap int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	fmt.Printf("%s: Склад открыт, вместимость %d\n", name, cap)
	defer fmt.Printf("%s: Склад закрыт, всего %d\n", name, count)

	for brick := range storageIn {
		if count < cap {
			count++
			fmt.Printf("%s: Принят %s кирпич, всего %d\n", name, brick.Color, count)
			time.Sleep(time.Millisecond * 100)
		} else {
			fmt.Printf("%s: Склад переполнен, кирпич потерян!\n", name)
		}
	}
}