/*
	У вас есть переменная r, которая содержит входные пользовательские данные.
	r – радиус.
	Напишите код, который находит периметр круга и записывает результат в переменную result.
	Формула расчета периметра круга: C=2⋅π⋅r , где C это периметр,r – радиус, π – число пи.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculateCircumference(r float64) float64 {
	return r * 2 * math.Pi
}

func main() {
	r := ReadInput()
	result := calculateCircumference(r)
	fmt.Println(fmt.Sprintf("%.2f", result))
}

func ReadInput() float64 {
	var r float64
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	r, _ = strconv.ParseFloat(input, 64)
	return r
}
