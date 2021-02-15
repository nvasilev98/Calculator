package calculator

func NewCalculatorError() CalculatorError {
	return CalculatorError{
		message: "Failed to calculate expression due to bad input",
	}
}

type CalculatorError struct {
	message string
}

func (e CalculatorError) Error() string {
	return e.message
}

func NewInvalidIDError() InvalidIDError {
	return InvalidIDError{
		message: "ID has not been found",
	}
}

type InvalidIDError struct {
	message string
}

func (e InvalidIDError) Error() string {
	return e.message
}

func NewEvaluationError() EvaluationError {
	return EvaluationError{
		message: "ID has not been evaluated",
	}
}

type EvaluationError struct {
	message string
}

func (e EvaluationError) Error() string {
	return e.message
}
