package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ROM2ARABIC map[string]int = map[string]int{
	"nulla": 0, "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11,
	"XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
}

var ARABIC2ROME map[int]string = map[int]string{
	0: "nulla", 1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
}

var OPERATORS []string = []string{
	"+", "-", "*", "/",
}

func contains(val string, arr []string) bool {
	result := false

	for _, value := range arr {
		if val == value {
			result = true
			break
		}
	}

	return result
}

func isOperator(operator string) bool {
	result := false
	if contains(operator, OPERATORS) {
		result = true
	}
	return result
}

func Convert2RomNumber(number int) string {
	return ARABIC2ROME[number]
}

func Convert2ArabicNumber(romNum string) int {
	return ROM2ARABIC[romNum]
}

func isRomeNumber(number string) bool {
	result := true
	_, ok := ROM2ARABIC[number]
	if !ok {
		result = false
	}
	return result
}

func isArabicNumber(number string) bool {
	result := true
	numberLI, err := strconv.Atoi(number)
	if err != nil {
		panic(fmt.Sprintf("Expression entered incorrectly: number %v entered incorrectly", number))
	} else {
		_, ok := ARABIC2ROME[numberLI]
		if !ok {
			result = false
		}
	}
	return result
}

func MakeCalc(val1 int, val2 int, op string) int {
	result := 0

	if val1 > 10 || val2 > 10 {
		panic("Expression entered incorrectly: result number is too big")
	}

	switch op {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		result = val1 / val2
	}

	return result
}

func CheckInputForNumberOfElements(input []string) {

	if len(input) != 3 {
		panic("Expression entered incorrectly: wrong number of arguments.")
	}

}

func ExpressionCheck(operator string, leftOperand string, rightOperand string, pred func(string) bool) {

	if !isOperator(operator) {
		panic(fmt.Sprintf("Expression entered incorrectly: operator %v entered incorrectly", operator))
	}
	if !pred(leftOperand) {
		panic(fmt.Sprintf("Expression entered incorrectly: number %v entered incorrectly", leftOperand))
	}
	if !pred(rightOperand) {
		panic(fmt.Sprintf("Expression entered incorrectly: number %v entered incorrectly", rightOperand))
	}

}

func CalcRome(left string, right string, operator string) string {
	result := 0
	ExpressionCheck(operator, left, right, isRomeNumber)
	result = MakeCalc(Convert2ArabicNumber(left), Convert2ArabicNumber(right), operator)
	if result < 0 {
		panic("The result of expression is a number less than zero. In Roman calculus there are no numbers less than zero.")
	}
	return Convert2RomNumber(result)
}

func CalcArabic(left string, right string, operator string) string {
	result := 0
	ExpressionCheck(operator, left, right, isArabicNumber)
	l, _ := strconv.Atoi(left)
	r, _ := strconv.Atoi(right)
	result = MakeCalc(l, r, operator)
	return strconv.Itoa(result)
}

func Calc(input []string) string {

	CheckInputForNumberOfElements(input)
	operator := input[1]
	leftOperand := input[0]
	rightOperand := input[2]

	var result string

	if isRomeNumber(leftOperand) {
		result = CalcRome(leftOperand, rightOperand, operator)
	} else {
		result = CalcArabic(leftOperand, rightOperand, operator)
	}

	return result
}

func SafeCalc() {
}

func main() { // не успел разобраться с обработкой ошибок в го в связи с чем закрываю программу через панику что не очень хорошо

	fmt.Print("WARNING: All elements of the expression must be entered with a space\n")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Type your expression: ")
		text, _, _ := reader.ReadLine()
		toCalc := strings.Split(string(text), " ")
		fmt.Print("Answer: " + Calc(toCalc) + "\n")
	}

}
