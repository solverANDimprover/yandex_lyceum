package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func MakeOperation(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("division by zero")
		}
		return a / b
	}
	return 0
}

func ParseNums(in string) ([]float64, error) {
	re := regexp.MustCompile(`\d+\.?\d*`)
	SliceOfMatches := re.FindAllString(in, -1)
	var NumsSlice []float64
	for i := 0; i < len(SliceOfMatches); i++ {
		intV, err := strconv.ParseFloat(SliceOfMatches[i], 64)
		if err != nil {
			panic(err)
		}
		NumsSlice = append(NumsSlice, intV)
	}
	return NumsSlice, nil
}

func CountMul(in string) int {
	re := regexp.MustCompile(`[/\*]`)
	SliceOfMatchesMul := re.FindAllString(in, -1)
	return len(SliceOfMatchesMul)
}

func ParseAdd(in string) []string {
	re := regexp.MustCompile(`[\+\-]`)
	var SliceOfAdd []string = []string{}
	SliceOfMatchesAdd := re.FindAllString(in, -1)
	for i := 0; i < len(SliceOfMatchesAdd); i++ {
		SliceOfAdd[i] = SliceOfMatchesAdd[i]
	}
	return SliceOfAdd
}

func ParseOps(in string) []string {
	re := regexp.MustCompile(`[\+\-\*/]`)
	SliceOfMatches := re.FindAllString(in, -1)
	return SliceOfMatches
}

func Calc(expression string) (float64, error) {
	Counter := 0
	for i := range expression {
		if string(expression[i]) == "(" || string(expression[i]) == ")" {
			Counter++
		}
	}
	if Counter%2 == 1 {
		return 0, fmt.Errorf("no")
	}
	Counter /= 2
	re := regexp.MustCompile(`\([0-9\.\+\*\-%^]*\)`)
	for i := 0; i < Counter; i++ {
		Match := re.FindString(expression)
		Match1 := Match[1 : len(Match)-1]
		Value, err := ParseValue(Match1)
		if err != nil {
			return 0, fmt.Errorf("panic")
		}
		expression = strings.Replace(expression, Match, Value, 1)
	}
	Val, err := ParseValue(expression)
	if err != nil {
		return 0, err
	}
	Answer, err := strconv.ParseFloat(Val, 64)
	if err != nil {
		return 0, err
	}
	return Answer, nil

}
func ParseValue(expression string) (string, error) {
	Nums, err := ParseNums(expression)
	Ops := ParseOps(expression)
	CounterMul := CountMul(expression)
	if len(Ops) >= len(Nums) {
		return "", fmt.Errorf("ops too much")
	}
	if err != nil {
		panic(err)
	}
	for i := 0; CounterMul > 0; i++ {
		if Ops[i] == "/" || Ops[i] == "*" {
			num1 := Nums[i]
			num2 := Nums[i+1]
			pt1 := append(Nums[:i], MakeOperation(num1, num2, Ops[i]))
			pt2 := append(pt1, Nums[i+2:]...)
			Nums = pt2
			Ops = append(Ops[:i], Ops[i+1:]...)
			CounterMul--
			i--
		}
	}
	for i := 0; len(Ops) > 0; i++ {
		num1 := Nums[i]
		num2 := Nums[i+1]
		pt1 := append(Nums[:i], MakeOperation(num1, num2, Ops[i]))
		pt2 := append(pt1, Nums[i+2:]...)
		Nums = pt2
		Ops = append(Ops[:i], Ops[i+1:]...)
		i--
	}
	return fmt.Sprintf("%f", Nums[0]), nil
}
func main() {
	fmt.Println(Calc("1+1*"))
}
