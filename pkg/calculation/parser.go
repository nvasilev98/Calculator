package calculation

import (
	"strconv"
	"strings"
)

//go:generate mockgen --source=parser.go --destination mocks/mock_parser.go --package mocks
type OperationFactorier interface {
	GetOperation(str string) (Evaluator, error)
}

func NewParser(operationfactorier OperationFactorier) *Parser {
	return &Parser{
		operationfactorier: operationfactorier,
	}
}

type Parser struct {
	operationfactorier OperationFactorier
}

func (p *Parser) Parse(str string) ([]string, error) {
	var result []string
	i := 0
	for i < len(str) {
		if str[i] == '(' || str[i] == ')' || str[i] == ' ' {
			if str[i] != ' ' {
				result = append(result, string(str[i]))
			}
		} else if str[i] >= '0' && str[i] <= '9' {
			number, err := parseNumber(i, str)
			if err != nil {
				return nil, err
			}
			result = append(result, number)
			i = i + len(number) - 1
		} else if _, err := p.operationfactorier.GetOperation(string(str[i])); err == nil {
			result = append(result, string(str[i]))
		} else {
			return nil, NewErrInvalidSymbol()
		}
		i++
	}
	if len(result) == 0 {
		return nil, NewErrEmptyString()
	}
	if !p.checkOperationOrder(result) {
		return nil, NewErrOperationOrder()
	}
	return result, nil
}

func isPartOfNumber(symbol byte) bool {
	return (symbol >= '0' && symbol <= '9') || symbol == '.'
}

func parseNumber(i int, str string) (string, error) {
	var sb strings.Builder
	for i < len(str) && isPartOfNumber(str[i]) {
		sb.WriteByte(str[i])
		i++
	}

	if _, err := strconv.ParseFloat(sb.String(), 64); err != nil {
		return "", NewErrInvalidFloatingNumber()
	}

	return sb.String(), nil

}

func (p *Parser) checkOperationOrder(expr []string) bool {
	if _, err := p.operationfactorier.GetOperation(expr[len(expr)-1]); err == nil {
		return false
	}
	countElem := 0
	for _, elem := range expr {

		if elem == "(" || elem == ")" {
			continue
		} else if _, err := p.operationfactorier.GetOperation(elem); err != nil {
			countElem++
		} else {
			countElem--
		}
		if countElem < 0 {
			return false
		}
	}
	if countElem != 1 {
		return false
	}
	return true
}
