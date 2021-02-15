package operation

type Extract struct{}

//Evaluate return substraction of two numbers
func (*Extract) Evaluate(num1, num2 float64) (float64, error) {
	return num1 - num2, nil
}

//Weight returns weight of operation
func (*Extract) Weight() int {
	return 1
}
