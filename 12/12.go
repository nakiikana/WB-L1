//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func CreateStringMultitude(arr []string) []string {
	//Множество состоит из уникальных значений - нам подойдут ключи карты.
	checkRepeat := make(map[string]struct{})
	for _, v := range arr {
		if _, ok := checkRepeat[v]; !ok {
			checkRepeat[v] = struct{}{}
		}
	}
	//Возвращаем массив из ключей.
	return maps.Keys(checkRepeat)
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(CreateStringMultitude(arr))
}
