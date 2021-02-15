package calculation

import (
	"github.com/golang-collections/collections/stack"
)

func NewReverseNotation(operationfactorier OperationFactorier) *ReverseNotation {
	return &ReverseNotation{
		operationfactorier: operationfactorier,
	}
}

type ReverseNotation struct {
	operationfactorier OperationFactorier
}

//Convert returns a postfix notation
func (rn *ReverseNotation) Convert(pStr []string) ([]string, error) {
	opStack := stack.New()
	var result []string
	var bracketCount int
	for _, elem := range pStr {
		if bracketCount < 0 {
			return nil, NewErrMismatchedBrackets()
		}
		if elem == "(" {
			bracketCount++
			opStack.Push(elem)
		} else if elem == ")" {
			for opStack.Len() > 0 && opStack.Peek().(string) != "(" {
				result = append(result, opStack.Pop().(string))
			}
			if opStack.Len() > 0 && opStack.Peek().(string) == "(" {
				opStack.Pop()
			}
			bracketCount--
		} else if op1, err := rn.operationfactorier.GetOperation(elem); err == nil {
			rn.appendOperation(&result, op1, opStack)
			opStack.Push(elem)
		} else {
			result = append(result, elem)
		}
	}
	if bracketCount != 0 {
		return nil, NewErrMismatchedBrackets()
	}
	for opStack.Len() > 0 {
		result = append(result, opStack.Pop().(string))
	}
	return result, nil
}

func precede(op1, op2 Evaluator) bool {
	return op1.Weight() <= op2.Weight()
}

func (rn *ReverseNotation) appendOperation(result *[]string, op1 Evaluator, opStack *stack.Stack) {
	if opStack.Len() > 0 {
		tmp := opStack.Peek().(string)
		if op2, err := rn.operationfactorier.GetOperation(tmp); err == nil {
			for opStack.Len() > 0 && precede(op1, op2) && opStack.Peek().(string) != "(" {
				*result = append(*result, opStack.Pop().(string))
				if opStack.Len() > 0 && opStack.Peek().(string) != "(" {
					tmp = opStack.Peek().(string)
					op2, err = rn.operationfactorier.GetOperation(tmp)
				}
			}
		}
	}
}
