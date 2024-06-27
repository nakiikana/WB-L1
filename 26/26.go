// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// 	aabcd — false

package main

import (
	"fmt"
	"strings"
)

func CheckIfUnique(str string) bool {
	//Переводим все буквы в строчные для регистронезависимости.
	lowerStr := strings.ToLower(str)
	symbols := make(map[rune]int)
	for _, v := range lowerStr {
		//Подсчитываем появление каждого символа.
		symbols[v] += 1
		//В случае повторения одного из них сразу возвращаем false.
		if symbols[v] > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("abcd - %v\n", CheckIfUnique("abcd"))
	fmt.Printf("abCdefAaf - %v\n", CheckIfUnique("abCdefAaf"))
	fmt.Printf("aabcd - %v\n", CheckIfUnique("aabcd"))
}
