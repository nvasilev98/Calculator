package calculation

import (
	"github.com/nvasilev98/calculator/pkg/calculation/operation"
)

//go:generate mockgen --source=operationfactory.go --destination mocks/mock_operationfactory.go --package mocks

type Evaluator interface {
	Evaluate(num1, num2 float64) (float64, error)
	Weight() int
}

func NewOperationFactory() *OperationFactory {
	return &OperationFactory{}
}

type OperationFactory struct{}

func (e *OperationFactory) GetOperation(op string) (Evaluator, error) {
	switch op {
	case "+":
		return &operation.Add{}, nil
	case "*":
		return &operation.Multiply{}, nil
	case "-":
		return &operation.Extract{}, nil
	case "/":
		return &operation.Divide{}, nil
	default:
		return nil, NewErrUnsupportedOperation()
	}
}
