/*
	У вас есть переменные year, month, day, hour, minute, second, которые содержат входные пользовательские данные.
	Напишите код, который приводит дату к такому формату %A %d-%m-%Y, %H:%M:%S и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func formatDate(year, month, day, hour, minute, second int) string {
	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
	result := t.Format("Friday 01-01-2006, 15:04:05")
	return result
}

func main() {
	year, month, day, hour, minute, second := ReadInput()
	result := formatDate(year, month, day, hour, minute, second)
	fmt.Println(result)
}

func ReadInput() (int, int, int, int, int, int) {
	var year, month, day, hour, minute, second int
	var values []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	year, _ = strconv.Atoi(values[0])
	month, _ = strconv.Atoi(values[1])
	day, _ = strconv.Atoi(values[2])
	hour, _ = strconv.Atoi(values[3])
	minute, _ = strconv.Atoi(values[4])
	second, _ = strconv.Atoi(values[5])
	return year, month, day, hour, minute, second
}
