//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func ReverseWordsInString(str string) string {
	reversed := ""
	//Разобьем исходную строку на части с разделителем " ".
	parts := strings.Split(str, " ")
	//Сконкатенируем новую строку в обратном порядке.
	for i := len(parts) - 1; i >= 0; i-- {
		reversed += parts[i] + " "
	}
	return reversed
}

func main() {
	fmt.Println(ReverseWordsInString("snow dog sun"))
	fmt.Println(ReverseWordsInString("the sun is so far away and yet so bright"))
}
