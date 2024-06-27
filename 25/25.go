//Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func Sleep(t time.Duration) {
	//time.After() используется для создания канала, который отправляет значение через указанный промежуток времени. На указанны промежуток
	//времени выполнение программы приостанавливается.
	<-time.After(t)
}

func main() {
	fmt.Println("This line is printed right after the start - wait for the next one to come in 5 minutes")
	Sleep(5 * time.Second)
	fmt.Println("Here it is")
}
