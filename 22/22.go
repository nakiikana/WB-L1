//Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a,b, значение которых > 2^20

package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
)

func Calculate(a, b *big.Int, operator string) *big.Int {
	res := new(big.Int)
	switch operator {
	case "+":
		res.Add(a, b)
	case "-":
		res.Sub(a, b)
	case "/":
		res.Div(a, b)
	case "*":
		res.Mul(a, b)
	default:
		log.Fatalf("Unsupported operator: %s", operator)
	}
	return res
}

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("You should consequently enter a, an operator and b")
	}
	a, ok := new(big.Int).SetString(os.Args[1], 10)
	if !ok {
		log.Fatalln("Could not read the first operand")
	}
	b, ok := new(big.Int).SetString(os.Args[3], 10)
	if !ok {
		log.Fatalln("Could not read the second operand")
	}
	operand := os.Args[2]
	res := Calculate(a, b, operand)
	fmt.Println(res)
}
