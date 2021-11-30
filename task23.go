package main

import (
	"fmt"
	"os"
)

func main() {
	// Это один из вариантов удаления из слайса, вроде бы самый быстрый т.к
	// в нем просто ненужный нам элемент перезаписываем на последний
	// и обрезаем слайс, чтобы последнего элемента не было в нем
	// в таком случае нарушается порядок
	// но о нем в задаче не говорилось)))
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var i int
	fmt.Fscan(os.Stdin, &i)
	a[i] = a[len(a)- 1]
	a = a[:len(a)-1]
	fmt.Print(a)
}
