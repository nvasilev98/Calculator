package fixture

import (
	"os/exec"
	"strconv"

	. "github.com/onsi/gomega"
)

type CalculatorPageObject struct {
	path string
}

func NewCalculatorPageObject(path string) *CalculatorPageObject {
	return &CalculatorPageObject{
		path: path,
	}
}

func (o *CalculatorPageObject) CalculateSuccessfully(expr string) float64 {
	res := o.executeCmd(expr)
	number, err := strconv.ParseFloat(res, 64)
	Expect(err).ToNot(HaveOccurred())
	return number
}

func (o *CalculatorPageObject) CalculateFails(expr string) {
	res := o.executeCmd(expr)
	_, err := strconv.ParseFloat(res, 64)
	Expect(err).To(HaveOccurred())
}
func (o *CalculatorPageObject) executeCmd(expr string) string {
	cmd := o.path + " \"" + expr + "\""
	res, err := exec.Command("bash", "-c", cmd).Output()
	Expect(err).ToNot(HaveOccurred())
	parsedRes := string(res)[:len(string(res))-1]
	return parsedRes
}
