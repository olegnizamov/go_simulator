/*
	У вас есть переменные n, e, которые содержат входные пользовательские данные.
	Напишите код, который возводит число n в степень e и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func power(n, e float64) float64 {
	return math.Pow(n, e)
}

func main() {
	n, e := ReadInput()
	result := power(n, e)
	fmt.Println(fmt.Sprintf("%.1f", result))
}

func ReadInput() (float64, float64) {
	var n float64
	var e float64
	var values []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	n, _ = strconv.ParseFloat(values[0], 64)
	e, _ = strconv.ParseFloat(values[1], 64)
	return n, e
}
