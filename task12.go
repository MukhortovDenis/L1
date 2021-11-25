package main
 
import "fmt"

func main() {
    a1 := []string{"cat", "cat", "dog", "cat", "tree"}
    a2 := newSet(a1)
    fmt.Println(a2)
}

func newSet(a1 []string) []string {
    keys := make(map[string]bool)
    a2 := []string{}	
    for _, str := range a1 {
		//Получаем булевое значение, 
		//Eсли нет такого ключа str, то получаем false и проходим по условию
		//Если есть такой ключ str, то получаем true и не проходим по условию
        if value := keys[str]; !value {
			// Устанавливаем для этого ключа значение true,
			// Чтобы следующая  строка равная одной из предыдущих не проходила условие
            keys[str] = true
            a2 = append(a2, str)
        }
    }    
    return a2
}