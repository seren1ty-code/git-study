package feature1

import (
	"fmt"
	"sync"
	"time"
)

func Feature1(number int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10; i++ {
		fmt.Println("Hello Git №", i, "GoRoutine:", number)
		time.Sleep(300 * time.Millisecond)
	}
}
