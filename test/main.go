package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pot() {
	count := 0
	for i := 0; i < 100; i++ {
		time.Sleep(400 * time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		if (rand.Intn(3)) == 2 {
			fmt.Println("healed")
			count += 1
		} else {
			fmt.Println("//")
		}
	}
	fmt.Println("total healed : ", count, "/100")
}
