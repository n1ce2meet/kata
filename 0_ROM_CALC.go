package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ROMAN2ARABIC map[string]int = map[string]int{
	"C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1,
}

var ROMAN2ARABICMAPORDER []string = []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"} // к сожалению привычный ordermap в стандарте отсутсвует

var OPERATORS string = "+-/*"

func isOperator(operator string) bool {
	result := false
	if strings.Contains(OPERATORS, operator) {
		result = true
	}
	return result
}

func isRomanNumeral(number string) bool {
	result := true
	_, ok := ROMAN2ARABIC[number]
	if !ok {
		result = false
	}
	return result
}

func toRoman(number int) string {
	result := ""

	if number == 0 {
		result = "nulla"
	} else {
		for _, roman := range ROMAN2ARABICMAPORDER[:] {
			arabic := ROMAN2ARABIC[roman]
			result += strings.Repeat(roman, (number / arabic))
			number %= arabic
		}
	}

	return result
}

func toArabic(number string) int {
	result := 0
	byNumeral := strings.Split(number, "")

	for i := 0; i < len(byNumeral); i++ {
		if isRomanNumeral(byNumeral[i]) {

			if i+1 < len(number) {
				if ROMAN2ARABIC[byNumeral[i]] < ROMAN2ARABIC[byNumeral[i+1]] {
					result += ROMAN2ARABIC[byNumeral[i]+byNumeral[i+1]]
					i++
					continue
				}
			}

			result += ROMAN2ARABIC[byNumeral[i]]

		} else {
			panic(fmt.Sprintf("Expression entered incorrectly: %v is not roman number", number))
		}
	}

	return result
}

func CheckExp(input []string) {

	if len(input) != 3 {
		panic("Expression entered incorrectly: wrong number of arguments.")
	}

	if !isOperator(input[1]) {
		panic("Expression entered incorrectly: operator entered incorrectly.")
	}

}

func CheckEnteredNumbers(numbers ...int) {
	for _, num := range numbers[:] {
		if num < 1 {
			panic(fmt.Sprintf("Expression entered incorrectly: number %v is too small", num))
		}
		if num > 10 {
			panic(fmt.Sprintf("Expression entered incorrectly: number %v is too big", num))
		}
	}
}

func MakeCalc(left int, right int, op string) int {
	result := 0

	CheckEnteredNumbers(left, right)

	switch op {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	}

	return result
}

func CalcRoman(left string, right string, operator string) string {
	result := 0
	result = MakeCalc(toArabic(left), toArabic(right), operator)
	if result < 0 {
		panic("The result of expression is a number less than zero. In Roman calculus there are no numbers less than zero.")
	}
	return toRoman(result)
}

func CalcArabic(left string, right string, operator string) string {
	result := 0
	l, lerr := strconv.Atoi(left)
	r, rerr := strconv.Atoi(right)
	if lerr != nil {
		panic(fmt.Sprintf("Expression entered incorrectly: number %v is incorrect", left))
	} else if rerr != nil {
		panic(fmt.Sprintf("Expression entered incorrectly: number %v is incorrect", right))
	}
	result = MakeCalc(l, r, operator)
	return strconv.Itoa(result)
}

func Calc(input []string) string {

	CheckExp(input)

	operator := input[1]
	leftOperand := input[0]
	rightOperand := input[2]

	var result string

	if isRomanNumeral(string(leftOperand[0])) {
		result = CalcRoman(leftOperand, rightOperand, operator)
	} else {
		result = CalcArabic(leftOperand, rightOperand, operator)
	}

	return result
}

func main() { // не успел разобраться с обработкой ошибок в го в связи с чем закрываю программу через панику что не очень хорошо

	fmt.Print("WARNING: All elements of the expression must be entered with a space\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type your expression: ")
	text, _, _ := reader.ReadLine()
	toCalc := strings.Split(string(text), " ")
	fmt.Print("Answer: " + Calc(toCalc) + "\n")

}
