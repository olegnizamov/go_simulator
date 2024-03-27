/*
	У вас есть переменная time, которая содержит входные пользовательские данные.
	Напишите код, переводит время из одного формата в другой и записывает результат в переменную result.
	Входящий формат: 12-часовой формат времени, где AM и PM — это сокращения, использующиеся для обозначения времени суток.
	например 9:49 AM перевести в 09:49
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toMilitaryTime(time string) string {
	period := time[len(time)-2:]
	timeParts := strings.Split(time[:len(time)-3], ":")
	hours := timeParts[0]
	minutes := timeParts[1]
	if period == "AM" {
		if len(hours) < 2 {
			hours = "0" + hours
		}
		return fmt.Sprintf("%s:%s", hours, minutes)
	} else {
		hoursInt, _ := strconv.Atoi(hours)
		hoursInt = hoursInt + 12
		hours = strconv.Itoa(hoursInt)
		return fmt.Sprintf("%s:%s", hours, minutes)
	}
}

func main() {
	time := ReadInput()
	result := toMilitaryTime(time)
	fmt.Println(result)
}

func ReadInput() string {
	var time string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	time = input
	return time
}
