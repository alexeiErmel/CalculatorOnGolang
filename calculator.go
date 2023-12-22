package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите операцию: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Неправильный формат введенных данных")
			continue
		}
		numeralsOne := parts[0]
		operation := parts[1]
		numeralsTwo := parts[2]
		answer := numeralsRoman(numeralsOne, operation, numeralsTwo)
		if answer == "Ошибка" {
			fmt.Println("Неправильный формат введенных данных")
			continue
		} else if answer == "Не римские цифры" {
			if numeralsArabic(numeralsOne, operation, numeralsTwo) == 100 {
				fmt.Println("Неправильный формат введенных данных")
				continue
			} else {
				fmt.Println(numeralsArabic(numeralsOne, operation, numeralsTwo))
			}
		} else {
			fmt.Println(answer)
		}
	}
}
func numeralsRoman(numeralsOne, operation, numeralsTwo string) string {
	romanNumerals := [21]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX"}
	var x, y int
	for i := 1; i <= 10; i++ {
		if romanNumerals[i] == numeralsOne {
			x = i
			break
		}
	}
	for i := 1; i <= 10; i++ {
		if romanNumerals[i] == numeralsTwo {
			y = i
			break
		}
	}
	if x != 0 && y != 0 {
		switch operation {
		case "+":
			return romanNumerals[amount(x, y)]
		case "-":
			return romanNumerals[subtraction(x, y)]
		case "*":
			return romanNumerals[multiply(x, y)]
		case "/":
			return romanNumerals[division(x, y)]
		}
	} else if x == 0 && y == 0 {
		return "Не римские цифры"
	} else {
		return "Ошибка"
	}
	return ""
}
func numeralsArabic(x, operation, y string) int {
	a, _ := strconv.Atoi(x)
	b, _ := strconv.Atoi(y)
	if a <= 10 && b <= 10 {
		switch operation {
		case "+":
			return amount(a, b)
		case "-":
			return subtraction(a, b)
		case "*":
			return multiply(a, b)
		case "/":
			return division(a, b)
		}
	} else {
		return 100
	}
	return 100
}
func amount(x, y int) int {
	return x + y
}
func subtraction(x, y int) int {
	return x - y
}
func multiply(x, y int) int {
	return x * y
}
func division(x, y int) int {
	return x / y
}
