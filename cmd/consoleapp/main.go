package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/nvasilev98/calculator/pkg/calculation"
	"github.com/nvasilev98/calculator/pkg/calculation/operation"
)

func checkArgumentsNumber(args []string) error {
	if len(args) != 2 {
		return errors.New("Not matching number of arguments")
	}
	return nil
}

func calculate(expr string) {
	var operationfactory = calculation.NewOperationFactory()
	var reader = calculation.NewParser(operationfactory)
	var converter = calculation.NewReverseNotation(operationfactory)
	calculator := calculation.NewCalculator(operationfactory, reader, converter)
	res, err := calculator.Calculate(expr)
	switch err.(type) {
	case operation.ErrZeroDivision:
		fmt.Println("Undefined behaviour")
	case calculation.ErrEmptyString:
		fmt.Println("Parsing failed due to empty input")
	case calculation.ErrInvalidFloatingNumber:
		fmt.Println("Parsing failed due to invalid floating")
	case calculation.ErrInvalidSymbol:
		fmt.Println("Parsing failed due to invalid symbol")
	case calculation.ErrOperationOrder:
		fmt.Println("Parsing failed due to invalid operation order")
	case calculation.ErrMismatchedBrackets:
		fmt.Println("Converting failed due to mismatched brakes")
	default:
		fmt.Println(res)
	}

}

func main() {
	args := os.Args
	if err := checkArgumentsNumber(args); err != nil {
		fmt.Println(err)
	} else {
		calculate(args[1])
	}
}
