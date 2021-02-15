package calculation

func NewErrUnsupportedOperation() ErrUnsupportedOperation {
	return ErrUnsupportedOperation{
		message: "Unsupported operation",
	}
}

type ErrUnsupportedOperation struct {
	message string
}

func (e ErrUnsupportedOperation) Error() string {
	return e.message
}
func NewErrInvalidSymbol() ErrInvalidSymbol {
	return ErrInvalidSymbol{
		message: "Invalid symbol has been found",
	}
}

type ErrInvalidSymbol struct {
	message string
}

func (e ErrInvalidSymbol) Error() string {
	return e.message
}

func NewErrEmptyString() ErrEmptyString {
	return ErrEmptyString{
		message: "Input is empty string",
	}
}

type ErrEmptyString struct {
	message string
}

func (e ErrEmptyString) Error() string {
	return e.message
}

func NewErrOperationOrder() ErrOperationOrder {
	return ErrOperationOrder{
		message: "Input has error with operations",
	}
}

type ErrOperationOrder struct {
	message string
}

func (e ErrOperationOrder) Error() string {
	return e.message
}

func NewErrInvalidFloatingNumber() ErrInvalidFloatingNumber {
	return ErrInvalidFloatingNumber{
		message: "Invalid floating number",
	}
}

type ErrInvalidFloatingNumber struct {
	message string
}

func (e ErrInvalidFloatingNumber) Error() string {
	return e.message
}

func NewErrMismatchedBrackets() ErrMismatchedBrackets {
	return ErrMismatchedBrackets{
		message: "Mismatched brackets",
	}
}

type ErrMismatchedBrackets struct {
	message string
}

func (e ErrMismatchedBrackets) Error() string {
	return e.message
}
