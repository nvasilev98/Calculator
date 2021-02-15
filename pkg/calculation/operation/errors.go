package operation

func NewErrZeroDivision() ErrZeroDivision {
	return ErrZeroDivision{
		message: "Division by 0",
	}
}

type ErrZeroDivision struct {
	message string
}

func (e ErrZeroDivision) Error() string {
	return e.message
}
