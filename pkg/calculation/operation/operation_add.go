package operation

type Add struct{}

//Evaluate returns sum of two numbers
func (*Add) Evaluate(num1, num2 float64) (float64, error) {
	return num1 + num2, nil
}

//Weight returns weight of operation
func (*Add) Weight() int {
	return 1
}
