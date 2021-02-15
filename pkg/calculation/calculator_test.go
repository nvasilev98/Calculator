package calculation_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/nvasilev98/calculator/pkg/calculation"
	. "github.com/nvasilev98/calculator/pkg/calculation/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator Unit Tests", func() {
	var (
		mockCtrl      *gomock.Controller
		mockReader    *MockReader
		mockConverter *MockConverter
		mockOperation *MockOperationFactorier
		mockEvaluator *MockEvaluator
		calc          *Calculator
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockReader = NewMockReader(mockCtrl)
		mockConverter = NewMockConverter(mockCtrl)
		mockEvaluator = NewMockEvaluator(mockCtrl)
		mockOperation = NewMockOperationFactorier(mockCtrl)
		calc = NewCalculator(mockOperation, mockReader, mockConverter)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Describe("Calculate Expression", func() {
		It("propagates error when validation fails", func() {
			mockReader.EXPECT().Parse("1+!").Return(nil, errors.New("Validation failed"))
			_, err := calc.Calculate("1+!")
			Expect(err).To(HaveOccurred())
		})
		It("propagates error when converter fails", func() {
			gomock.InOrder(
				mockReader.EXPECT().Parse("1+2)").Return([]string{"1", "+", "2", ")"}, nil),
				mockConverter.EXPECT().Convert([]string{"1", "+", "2", ")"}).Return(nil, errors.New("Conversion failed")),
			)
			_, err := calc.Calculate("1+2)")
			Expect(err).To(HaveOccurred())
		})
		It("propagates error when operation fails", func() {
			gomock.InOrder(
				mockReader.EXPECT().Parse("11/0").Return([]string{"11", "/", "0"}, nil),
				mockConverter.EXPECT().Convert([]string{"11", "/", "0"}).Return([]string{"11", "0", "/"}, nil),
				mockOperation.EXPECT().GetOperation("11").Return(nil, errors.New("Number1 was picked")),
				mockOperation.EXPECT().GetOperation("0").Return(nil, errors.New("Number2 was picked")),
				mockOperation.EXPECT().GetOperation("/").Return(mockEvaluator, nil),
				mockEvaluator.EXPECT().Evaluate(float64(11), float64(0)).Return(float64(0), errors.New("Division by zero")),
			)
			_, err := calc.Calculate("11/0")
			Expect(err).To(HaveOccurred())
		})
		It("calculates single operation expression", func() {
			gomock.InOrder(
				mockReader.EXPECT().Parse("11+10").Return([]string{"11", "+", "10"}, nil),
				mockConverter.EXPECT().Convert([]string{"11", "+", "10"}).Return([]string{"11", "10", "+"}, nil),
				mockOperation.EXPECT().GetOperation("11").Return(nil, errors.New("Number1 was picked")),
				mockOperation.EXPECT().GetOperation("10").Return(nil, errors.New("Number2 was picked")),
				mockOperation.EXPECT().GetOperation("+").Return(mockEvaluator, nil),
				mockEvaluator.EXPECT().Evaluate(float64(11), float64(10)).Return(float64(21), nil),
			)
			res, err := calc.Calculate("11+10")
			Expect(err).To(BeNil())
			Expect(res).To(Equal(float64(21)))
		})
		It("calculates without expression parentheses", func() {
			gomock.InOrder(
				mockReader.EXPECT().Parse("11+3/2").Return([]string{"11", "+", "3", "/", "2"}, nil),
				mockConverter.EXPECT().Convert([]string{"11", "+", "3", "/", "2"}).Return([]string{"11", "3", "2", "/", "+"}, nil),
				mockOperation.EXPECT().GetOperation("11").Return(nil, errors.New("Number1 was picked")),
				mockOperation.EXPECT().GetOperation("3").Return(nil, errors.New("Number2 was picked")),
				mockOperation.EXPECT().GetOperation("2").Return(nil, errors.New("Number3 was picked")),
				mockOperation.EXPECT().GetOperation("/").Return(mockEvaluator, nil),
				mockEvaluator.EXPECT().Evaluate(float64(3), float64(2)).Return(float64(1.5), nil),
				mockOperation.EXPECT().GetOperation("+").Return(mockEvaluator, nil),
				mockEvaluator.EXPECT().Evaluate(float64(11), float64(1.5)).Return(float64(12.5), nil),
			)
			res, err := calc.Calculate("11+3/2")
			Expect(err).To(BeNil())
			Expect(res).To(Equal(float64(12.5)))
		})
		It("calculates with expression parentheses", func() {
			gomock.InOrder(
				mockReader.EXPECT().Parse("(11+3)/2").Return([]string{"(", "11", "+", "3", ")", "/", "2"}, nil),
				mockConverter.EXPECT().Convert([]string{"(", "11", "+", "3", ")", "/", "2"}).Return([]string{"11", "3", "+", "2", "/"}, nil),
				mockOperation.EXPECT().GetOperation("11").Return(nil, errors.New("Number1 was picked")),
				mockOperation.EXPECT().GetOperation("3").Return(nil, errors.New("Number2 was picked")),
				mockOperation.EXPECT().GetOperation("+").Return(mockEvaluator, nil),
				mockEvaluator.EXPECT().Evaluate(float64(11), float64(3)).Return(float64(14), nil),
				mockOperation.EXPECT().GetOperation("2").Return(nil, errors.New("Number3 was picked")),
				mockOperation.EXPECT().GetOperation("/").Return(mockEvaluator, nil),
				mockEvaluator.EXPECT().Evaluate(float64(14), float64(2)).Return(float64(7), nil),
			)
			res, err := calc.Calculate("(11+3)/2")
			Expect(err).To(BeNil())
			Expect(res).To(Equal(float64(7)))
		})
	})
})
