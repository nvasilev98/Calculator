package calculation

import (
	"strconv"

	"github.com/golang-collections/collections/stack"
)

//go:generate mockgen --source=calculator.go --destination mocks/mock_calculator.go --package mocks

type Reader interface {
	Parse(expr string) ([]string, error)
}

type Converter interface {
	Convert(expr []string) ([]string, error)
}

func NewCalculator(operationfactorier OperationFactorier, reader Reader, converter Converter) *Calculator {
	return &Calculator{
		operationfactorier: operationfactorier,
		reader:             reader,
		converter:          converter,
	}
}

type Calculator struct {
	operationfactorier OperationFactorier
	reader             Reader
	converter          Converter
}

func (c *Calculator) Calculate(expr string) (float64, error) {
	res, err := c.reader.Parse(expr)
	if err != nil {
		return 0, err
	}
	rpnExpr, err := c.converter.Convert(res)
	if err != nil {
		return 0, err
	}
	evStack := stack.New()
	for _, elem := range rpnExpr {
		if op, isOperation := c.determineElement(elem); isOperation == true {
			num2 := evStack.Pop().(float64)
			num1 := evStack.Pop().(float64)
			res, err := op.Evaluate(num1, num2)
			if err != nil {
				return 0, err
			}
			evStack.Push(res)
		} else {
			if num, err := strconv.ParseFloat(elem, 64); err == nil {
				evStack.Push(num)
			}
		}
	}
	return evStack.Peek().(float64), nil
}

func (c *Calculator) determineElement(elem string) (Evaluator, bool) {
	operation, err := c.operationfactorier.GetOperation(elem)
	if err != nil {
		return nil, false
	}
	return operation, true
}
