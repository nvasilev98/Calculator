package operation

type Multiply struct{}

//Evaluate returns product of two numbers
func (*Multiply) Evaluate(num1, num2 float64) (float64, error) {
	return num1 * num2, nil
}

//Weight returns weight of operation
func (*Multiply) Weight() int {
	return 2
}
