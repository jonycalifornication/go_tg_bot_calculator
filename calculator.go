package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func calculator(s string) string {
	expression, err := govaluate.NewEvaluableExpression(s)
	if err != nil {
		return "Invalid expression"
	}
	print("expression:", expression)

	result, err := expression.Evaluate(nil)
	if err != nil {
		return "Invalid expression"
	}
	print("result:", result)

	return fmt.Sprintf("%d", result)
}
