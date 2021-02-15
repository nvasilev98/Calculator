package operation

type Divide struct{}

//Evaluate returns division of two numbers
func (*Divide) Evaluate(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, NewErrZeroDivision()
	}
	return num1 / num2, nil
}

//Weight returns weight of operation
func (*Divide) Weight() int {
	return 2
}
