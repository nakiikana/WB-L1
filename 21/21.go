//Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// Создаем структуру, которую будем адаптировать под определенный интерфейс (target).
type Appartment struct {
	inhabitants int
	floor       int
}

// Добавим для данной структуры ее собственные методы.
func (a *Appartment) IsOnTheFloor() {
	fmt.Printf("The appartment is located on the %v floor\n", a.floor)
}

func (a *Appartment) HasInhabitants() {
	fmt.Printf("There are %v people living in the appartment\n", a.inhabitants)
}

// Интерфейс, под который мы будем адаптировать структуру Appartment (target).
type RealEstateAdapter interface {
	IsLocated()
}

// Сам адаптер с внедрением Appartment.
type AppartmentAdapter struct {
	*Appartment
}

// Конструктор для адаптера.
func NewAppartmentAdapter(a *Appartment) *AppartmentAdapter {
	return &AppartmentAdapter{a}
}

// Теперь адаптер удовлетворяет целевому интерфейсу.
func (a *AppartmentAdapter) IsLocated() {
	a.IsOnTheFloor()
}

// Создадим также для наглядности структуру, которая изначально удовлетворяет целевому интерфейсу:
type Garage struct{}

func (g *Garage) IsLocated() {
	fmt.Printf("Your garage is located in your yard\n")
}

func main() {
	//Демонстрация работы.
	everythingIOwn := [2]RealEstateAdapter{NewAppartmentAdapter(&Appartment{inhabitants: 3, floor: 13}), &Garage{}}
	for _, myRealEstate := range everythingIOwn {
		myRealEstate.IsLocated()
	}
}
