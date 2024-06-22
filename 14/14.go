//Разработать программу, которая в рантайме способна определить
//тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

// С помощью пакета reflect возможно определить тип переменной в рантайне:
func defineType(x interface{}) {
	xType := reflect.TypeOf(x)
	xValue := reflect.ValueOf(x)
	fmt.Println("The type of ", xValue, " is ", xType)
}

func main() {
	i := 5
	defineType(i)

	str := "cat"
	defineType(str)

	b := true
	defineType(b)

	ch := make(chan bool)
	defineType(ch)
}
