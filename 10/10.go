//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов.
//Последовательность в подмножноствах не важна.

package main

import (
	"fmt"
	"math"
)

func MakeTemperatureGroup(tempArr []float64) map[int][]float64 {
	resultMap := make(map[int][]float64)
	for _, v := range tempArr {
		key := int(v - math.Mod(v, 10))
		if _, ok := resultMap[key]; !ok {
			resultMap[key] = []float64{}
		}
		resultMap[key] = append(resultMap[key], v)
	}
	return resultMap
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(MakeTemperatureGroup(arr))
}
