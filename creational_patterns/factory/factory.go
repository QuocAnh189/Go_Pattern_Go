package main

import "fmt"

type Operator interface {
	SetOperands(a, b int)
	Calculate() int
}

type PlusOperator struct {
	a, b int
}

type MinusOperator struct {
	a, b int
}

func (o *PlusOperator) SetOperands(a, b int) {
	o.a = a
	o.b = b
}

func (o *PlusOperator) Calculate() int {
	return o.a + o.b
}

func (o *MinusOperator) SetOperands(a, b int) {
	o.a = a
	o.b = b
}

func (o *MinusOperator) Calculate() int {
	return o.a - o.b
}

func OperatorFactory(operatorType string) Operator {
	switch operatorType {
	case "+":
		return &PlusOperator{}
	case "-":
		return &MinusOperator{}
	default:
		return nil
	}
}

func main() {
	addition := OperatorFactory("+")
	addition.SetOperands(10, 5)
	fmt.Println("10 + 5 =", addition.Calculate()) // 10 + 5 = 15

	subtraction := OperatorFactory("-")
	subtraction.SetOperands(10, 5)
	fmt.Println("10 - 5 =", subtraction.Calculate()) // 10 - 5 = 5
}
