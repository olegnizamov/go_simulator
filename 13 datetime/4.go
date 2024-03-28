/*
	У вас есть переменные start, end, которые содержат входные пользовательские данные.
	Напишите код, подсчитывает количество дней между двумя датами и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func diff(start, end string) int {
	dateStart, _ := time.Parse("2006-01-02", start)
	dateEnd, _ := time.Parse("2006-01-02", end)
	days := int(dateEnd.Sub(dateStart).Hours()/24) + 1
	return days
}

func main() {
	start, end := ReadInput()
	result := diff(start, end)
	fmt.Println(result)
}

func ReadInput() (string, string) {
	var start, end string
	var values []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	start = values[0]
	end = values[1]
	return start, end
}
