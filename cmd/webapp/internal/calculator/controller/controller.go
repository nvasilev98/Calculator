package controller

import (
	"github.com/nvasilev98/calculator/cmd/webapp/internal/calculator"
	"github.com/nvasilev98/calculator/pkg/database/postgres"
	"github.com/pkg/errors"
)

//go:generate mockgen --source=controller.go --destination mocks/mock_controller.go --package mocks
type Calculator interface {
	Calculate(expr string) (float64, error)
}
type DataAccess interface {
	InsertCalculation(model postgres.Model) error
	SelectResultByID(id int64) (postgres.Response, bool, error)
}

func NewController(calculator Calculator, dataaccess DataAccess) *Controller {
	return &Controller{
		calculator: calculator,
		dataaccess: dataaccess,
	}
}

type Controller struct {
	calculator Calculator
	dataaccess DataAccess
}

func (c *Controller) Calculate(expr string) (float64, error) {
	return c.calculator.Calculate(expr)
}
func (c *Controller) Insert(model postgres.Model) error {
	return c.dataaccess.InsertCalculation(model)
}

func (c *Controller) SelectResult(id int64) (float64, error) {
	response, found, err := c.dataaccess.SelectResultByID(id)
	if err != nil {
		return 0, errors.Wrap(err, "failed interaction with database")
	}
	if !found {
		return 0, calculator.NewInvalidIDError()
	}
	if response.Error == "" {
		return 0, calculator.NewEvaluationError()
	}
	if response.Error != "nil" {
		return 0, calculator.NewCalculatorError()
	}

	return response.Result, nil
}
