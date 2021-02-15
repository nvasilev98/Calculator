package calculation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/nvasilev98/calculator/pkg/calculation"
)

var _ = Describe("Reverse Notation Unit Tests", func() {
	var (
		operationfactory = NewOperationFactory()
		reverse          = NewReverseNotation(operationfactory)
	)
	Describe("Reverse Expression", func() {
		It("reverses expression with one operation correctly", func() {
			res, err := reverse.Convert([]string{"3"})
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"3"}))
		})
		It("reverses expression with one operand and with brackets correctly", func() {
			res, err := reverse.Convert([]string{"(", "3.5", ")"})
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"3.5"}))
		})
		It("returns error if only opening bracket", func() {
			_, err := reverse.Convert([]string{"(", "3"})
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(ErrMismatchedBrackets{}))
		})
		It("returns error if mismatched bracket ", func() {
			_, err := reverse.Convert([]string{")", "3", "("})
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(ErrMismatchedBrackets{}))
		})
		It("returns error if only closing bracket", func() {
			_, err := reverse.Convert([]string{"3", ")"})
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(ErrMismatchedBrackets{}))
		})
		It("reverses simple expression correctly", func() {
			res, err := reverse.Convert([]string{"3", "+", "5"})
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"3", "5", "+"}))
		})
		It("reverses simple expression with brackets correctly", func() {
			res, err := reverse.Convert([]string{"(", "3", "+", "5", ")"})
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"3", "5", "+"}))
		})
		It("reverses expression correctly", func() {
			res, err := reverse.Convert([]string{"3", "+", "5", "*", "2"})
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"3", "5", "2", "*", "+"}))
		})
		It("reverses complicated expression correctly", func() {
			res, err := reverse.Convert([]string{"3", "*", "2", "/", "2", "+", "13", "-", "1", "/", "4"})
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"3", "2", "*", "2", "/", "13", "+", "1", "4", "/", "-"}))
		})
		It("reverses expression with brackets before operation correctly", func() {
			res, err := reverse.Convert([]string{"(", "3", "+", "5", ")", "*", "4"})
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"3", "5", "+", "4", "*"}))
		})
		It("reverses expression with brackets after operation correctly", func() {
			res, err := reverse.Convert([]string{"3.3", "*", "(", "5", "+", "4", ")"})
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"3.3", "5", "4", "+", "*"}))
		})
		It("reverses complicated expression with brackets correctly", func() {
			res, err := reverse.Convert([]string{"(", "2", "+", "4", "*", "3", "/", "1", ")"})
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"2", "4", "3", "*", "1", "/", "+"}))
		})

	})
})
