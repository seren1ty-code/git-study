package main

import (
	"fmt"
	"math/rand"
	feature1 "study/Feature1"
	"study/feature2"
	"sync"
)

func main() {
	defer fmt.Println("main func completed")

	var wg sync.WaitGroup
	minInt := 10
	fmt.Println("Hello, World!")
	for i := 1; i <= minInt+rand.Intn(15); i++ {
		wg.Add(1)
		go feature1.Feature1(i, &wg)
	}

	wg.Wait()

	feature2.Feature2()
}
