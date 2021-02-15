package controller_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/nvasilev98/calculator/cmd/webapp/internal/calculator"
	. "github.com/nvasilev98/calculator/cmd/webapp/internal/calculator/controller"
	. "github.com/nvasilev98/calculator/cmd/webapp/internal/calculator/controller/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nvasilev98/calculator/pkg/database/postgres"
)

var _ = Describe("Controller", func() {
	var (
		mockCtrl       *gomock.Controller
		mockCalculator *MockCalculator
		mockDataAccess *MockDataAccess
		controller     *Controller
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCalculator = NewMockCalculator(mockCtrl)
		mockDataAccess = NewMockDataAccess(mockCtrl)
		controller = NewController(mockCalculator, mockDataAccess)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("calculates single operation expression", func() {
		mockCalculator.EXPECT().Calculate("43+2").Return(float64(3), nil)
		res, err := controller.Calculate("43+2")
		Expect(err).ToNot(HaveOccurred())
		Expect(res).To(Equal(float64(3)))
	})
	It("inserts expression", func() {
		mockDataAccess.EXPECT().InsertCalculation(postgres.Model{Expression: "43+2"}).Return(nil)
		err := controller.Insert(postgres.Model{Expression: "43+2"})
		Expect(err).ToNot(HaveOccurred())
	})

	It("successfully select result with valid id", func() {
		mockDataAccess.EXPECT().SelectResultByID(int64(2)).Return(postgres.Response{Result: 5, Error: "nil"}, true, nil)
		result, err := controller.SelectResult(2)
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(float64(5)))
	})
	It("propagates error when interaction with database fail", func() {
		mockDataAccess.EXPECT().SelectResultByID(int64(1)).Return(postgres.Response{}, false, errors.New("failed to interact with database"))
		_, err := controller.SelectResult(1)
		Expect(err).To(HaveOccurred())
	})
	It("propagates error when invalid id is selected", func() {
		mockDataAccess.EXPECT().SelectResultByID(int64(1)).Return(postgres.Response{}, false, nil)
		_, err := controller.SelectResult(1)
		Expect(err).To(HaveOccurred())
		Expect(err).To(BeAssignableToTypeOf(calculator.InvalidIDError{}))
	})
	It("returns error when there is a calculation error", func() {
		mockDataAccess.EXPECT().SelectResultByID(int64(2)).Return(postgres.Response{Result: 0, Error: "Mismatched brakes"}, true, nil)
		_, err := controller.SelectResult(2)
		Expect(err).To(HaveOccurred())
		Expect(err).To(BeAssignableToTypeOf(calculator.CalculatorError{}))
	})
})
