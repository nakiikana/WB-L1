//Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type human struct {
	name    string
	surname string
	age     int
}

// ранее созданная структура human внедрена в action
type action struct {
	name  string
	speed float32
	human
}

func (h human) humanData() {
	fmt.Printf("%s %s - %d", h.name, h.surname, h.age)
}

// метод структуры human humanData встроен в action и доступен напрямаю
// (в случае существование одноименного метода у структуры action приоритет отдается ему)
func main() {
	sings := &action{name: "singing", speed: 3.12, human: human{name: "Jack", surname: "Forman", age: 34}}
	sings.humanData()
}
