package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	err              = errors.New("неправильный формат введенных данных")
	errRomanNegative = errors.New("римские числа не могут быть отрицательными,или равными нулю")
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите операцию: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("", err)
			continue
		}
		numeralsOne := parts[0]
		operation := parts[1]
		numeralsTwo := parts[2]
		answer, errAnswer := numeralsRoman(numeralsOne, operation, numeralsTwo)
		if errAnswer != nil {
			fmt.Println("", err)
			break
		} else {
			if answer == "Арабские цифры" {
				answerArabic, errArabic := numeralsArabic(numeralsOne, operation, numeralsTwo)
				if errArabic != nil {
					fmt.Println("", err)
					break
				} else {
					fmt.Println(answerArabic)
				}
			} else {
				fmt.Println(answer)
			}
		}
	}
}

func numeralsRoman(numeralsOne, operation, numeralsTwo string) (string, error) {
	romanNumerals := [101]string{
  "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
  "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
  "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
  "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
  "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
  "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
  "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
  "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
  "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI","LXXXVII","LXXXVIII","LXXXIX","XC",
   "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
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
		if x >= 1 && y >= 1 {
			switch operation {
			case "+":
				return romanNumerals[amount(x, y)], nil
			case "-":
				if subtraction(x, y) > 0 {
					return romanNumerals[subtraction(x, y)], nil
				} else {
					return "", errRomanNegative
				}
			case "*":
				return romanNumerals[multiply(x, y)], nil
			case "/":
				return romanNumerals[division(x, y)], nil
			}
		} else {
			return "", err
		}
	} else if x == 0 && y == 0 {
		return "Арабские цифры", nil
	} else {
		return "", err
	}
	return "", err
}
func numeralsArabic(x, operation, y string) (int, error) {
	a, errIntOne := strconv.Atoi(x)
	if errIntOne != nil {
		return 0, err
	}
	b, errIntTwo := strconv.Atoi(y)
	if errIntTwo != nil {
		return 0, err
	}
	if a <= 10 && b <= 10 {
		switch operation {
		case "+":
			return amount(a, b), nil
		case "-":
			return subtraction(a, b), nil
		case "*":
			return multiply(a, b), nil
		case "/":
			return division(a, b), nil
		}
	} else {
		return 0, err
	}
	return 0, err
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
