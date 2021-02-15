package calculation_test

import (
	. "github.com/nvasilev98/calculator/pkg/calculation"
	"github.com/nvasilev98/calculator/pkg/calculation/operation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OperationFactory Unit Tests", func() {
	var (
		operationfactory = NewOperationFactory()
	)
	It("returns error if operation is unsupported", func() {
		_, err := operationfactory.GetOperation("@")
		Expect(err).To(HaveOccurred())
		Expect(err).To(BeAssignableToTypeOf(ErrUnsupportedOperation{}))
	})
	It("returns correct operation to match multiplication", func() {
		mptly, err := operationfactory.GetOperation("*")
		Expect(err).To(BeNil())
		Expect(mptly).To(Equal(&operation.Multiply{}))
	})
	It("returns correct operation to match addition", func() {
		add, err := operationfactory.GetOperation("+")
		Expect(err).To(BeNil())
		Expect(add).To(Equal(&operation.Add{}))

	})
	It("returns correct operation to match substraction", func() {
		extract, err := operationfactory.GetOperation("-")
		Expect(err).To(BeNil())
		Expect(extract).To(Equal(&operation.Extract{}))
	})
	It("returns correct operation to match division", func() {
		divide, err := operationfactory.GetOperation("/")
		Expect(err).To(BeNil())
		Expect(divide).To(Equal(&operation.Divide{}))
	})
})
