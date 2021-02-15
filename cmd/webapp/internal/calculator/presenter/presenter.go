package presenter

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nvasilev98/calculator/cmd/webapp/internal/calculator"
	"github.com/nvasilev98/calculator/pkg/api"
	"github.com/nvasilev98/calculator/pkg/calculation"
	"github.com/nvasilev98/calculator/pkg/calculation/operation"
	"github.com/nvasilev98/calculator/pkg/database/postgres"
)

//go:generate mockgen --source=presenter.go --destination mocks/mock_presenter.go --package mocks

type Controller interface {
	Calculate(expr string) (float64, error)
	Insert(model postgres.Model) error
	SelectResult(id int64) (float64, error)
}

func NewPresenter(controller Controller) *Presenter {
	return &Presenter{
		controller: controller,
	}
}

type Presenter struct {
	controller Controller
}

func (p *Presenter) CalculateExpression(ctx *gin.Context) {
	var calc api.RequestBody
	err := ctx.ShouldBindJSON(&calc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		result, err := p.controller.Calculate(calc.Expression)
		wrappedErr := wrapError(err)
		if wrappedErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": wrappedErr.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"result": result})
		}
	}
}

func (p *Presenter) InsertExpression(ctx *gin.Context) {
	var calc api.RequestBody
	err := ctx.ShouldBindJSON(&calc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		err := p.controller.Insert(postgres.Model{Expression: calc.Expression})
		wrappedErr, statusCode := wrapErrorAndStatusCode(err)
		if wrappedErr != nil {
			ctx.JSON(statusCode, gin.H{"error": wrappedErr.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "successfully inserted row into database"})
		}
	}
}

func (p *Presenter) GetExpression(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		result, err := p.controller.SelectResult(intID)
		wrappedErr, statusCode := wrapErrorAndStatusCode(err)
		if wrappedErr != nil {
			ctx.JSON(statusCode, gin.H{"error": wrappedErr.Error()})
		} else {
			ctx.JSON(statusCode, gin.H{"result": result})
		}
	}
}

func wrapError(err error) error {
	if err != nil {
		switch err.(type) {
		case operation.ErrZeroDivision:
			return errors.New("Undefined behaviour")
		case calculation.ErrEmptyString:
			return errors.New("Parsing failed due to empty input")
		case calculation.ErrInvalidFloatingNumber:
			return errors.New("Parsing failed due to invalid floating")
		case calculation.ErrInvalidSymbol:
			return errors.New("Parsing failed due to invalid symbol")
		case calculation.ErrOperationOrder:
			return errors.New("Parsing failed due to invalid operation order")
		case calculation.ErrMismatchedBrackets:
			return errors.New("Converting failed due to mismatched brakes")
		}
	}
	return nil
}

func wrapErrorAndStatusCode(err error) (error, int) {
	if err != nil {
		switch err.(type) {
		case calculator.EvaluationError:
			return errors.New("Failed to return result due to evaluation is in progress"), http.StatusBadRequest
		case calculator.CalculatorError:
			return errors.New("Calculation failed due to bad input"), http.StatusOK
		case calculator.InvalidIDError:
			return errors.New("Failed to return result due to invalid id"), http.StatusBadRequest
		default:
			return errors.New("Failed due to internal error"), http.StatusInternalServerError
		}
	}
	return nil, http.StatusOK
}
