package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	var a int
	fmt.Println("Выберите способ остановки горутины(1 или 2)")
	fmt.Fscan(os.Stdin, &a)
	ch := make(chan struct{})
	defer close(ch)
	wg.Add(1)
	if a == 1 {

		go func() {
			for {
				select {
				case <-ch:
					fmt.Println("При прочтении из канала завершаем функцию с помощью return")
					wg.Done()
					return
				default:
					fmt.Println("Первая горутина работает")
				}
			}
		}()

		time.Sleep(1 * time.Second)
		ch <- struct{}{}
		wg.Wait()
		fmt.Println("Дождались завершения выбранной горутины")

	} else if a == 2 {

		go func() {
			for {
				select {
				case <-ch:
					fmt.Println("Вторая горутина работает")
				case ch <- struct{}{}: 
					wg.Done()
					fmt.Println("Даем сигнал в канал из горутины, после чего канал закрывается")
					return
				}
			}
		}()
		for i := 0; i < 10; i++{
			ch <- struct{}{}
		}
		<-ch
		wg.Wait()
		fmt.Println("Дождались завершения выбранной горутины")
	} else {
		fmt.Println("Ввели несуществующий вариант")
	}
}
