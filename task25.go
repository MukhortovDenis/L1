package main

import (
	"context"
	"fmt"
	"time"
)

func sleep(i time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*i)
	defer cancel()
	// Канал блокирует основную горутину
	<-ctx.Done()
	return

}

func main() {
	sleep(5)
	fmt.Println("Cон закончился")
}
