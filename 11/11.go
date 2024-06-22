//Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func FindIntersection(a []int, b []int) []int {
	intersection := []int{}
	//Так как поиск по карте оптимальнее, чем прохождение по всему множеству, приводим значения к соответствующему виду.
	//Значения по ключу нас не интересуют - их можно заменить на ничего не весящие пустые структуры.
	aMap := make(map[int]struct{})
	for _, v := range a {
		aMap[v] = struct{}{}
	}
	//Значения в множествах уникальны - нет необхоимости проводить дополнительные проверки пересечений.
	for _, v := range b {
		//Находим пересечения.
		if _, ok := aMap[v]; ok {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

func main() {
	a := []int{1, 17, 3, 23, 8, 11, 0}
	b := []int{3, 2, 11, 23, 4, 6, -2}
	fmt.Println(FindIntersection(a, b))
	c := []int{2, 3, 12, -4, 22}
	d := []int{-4, 33, -22, 3, 89}
	fmt.Println(FindIntersection(c, d))
}
