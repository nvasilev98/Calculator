package presenter_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/nvasilev98/calculator/cmd/webapp/internal/calculator"
	. "github.com/nvasilev98/calculator/cmd/webapp/internal/calculator/presenter"
	. "github.com/nvasilev98/calculator/cmd/webapp/internal/calculator/presenter/mocks"
	"github.com/nvasilev98/calculator/pkg/api"
	"github.com/nvasilev98/calculator/pkg/calculation"
	"github.com/nvasilev98/calculator/pkg/calculation/operation"
	"github.com/nvasilev98/calculator/pkg/database/postgres"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/resty.v1"
)

var _ = Describe("Presenter Unit Tests", func() {
	var gomockCtrl *gomock.Controller
	var httpServer *httptest.Server
	var req *resty.Request
	var router *gin.Engine
	var controller *MockController
	var presenter *Presenter
	var calcexpr api.Response
	var databaseResponse api.DatabaseResponse
	var databaseErrResponse api.DatabaseErrorResponse

	BeforeEach(func() {
		gomockCtrl, _ = gomock.WithContext(context.Background(), GinkgoT())
		router = gin.Default()
		httpServer = httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			defer GinkgoRecover()
			router.ServeHTTP(resp, req)
		}))
		req = resty.SetHostURL(httpServer.URL).R()
		controller = NewMockController(gomockCtrl)
		presenter = NewPresenter(controller)
		router.POST("/calculate", presenter.CalculateExpression)
		router.POST("/create", presenter.InsertExpression)
		router.GET("/get/:id", presenter.GetExpression)
		calcexpr = api.Response{}
	})
	AfterEach(func() {
		gomockCtrl.Finish()
	})
	It("returns bad request status code when invalid expression value", func() {
		resp, err := req.SetBody(`{"Expression": 1"3*2+1"}`).SetResult(&calcexpr).Post("/calculate")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusBadRequest))

	})
	It("returns correct sum and status code", func() {
		controller.EXPECT().Calculate("1+2").Return(float64(3), nil)
		resp, err := req.SetBody(api.RequestBody{Expression: "1+2"}).SetResult(&calcexpr).Post("/calculate")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		Expect(calcexpr.Result).To(Equal(float64(3)))
	})
	It("returns bad request status code when dividing by 0", func() {
		controller.EXPECT().Calculate("4-1/0").Return(float64(0), operation.NewErrZeroDivision())
		resp, err := req.SetBody(api.RequestBody{Expression: "4-1/0"}).SetResult(&calcexpr).Post("/calculate")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusBadRequest))

	})
	It("returns bad request status code when input is emtpy", func() {
		controller.EXPECT().Calculate(" ").Return(float64(0), calculation.NewErrEmptyString())
		resp, err := req.SetBody(api.RequestBody{Expression: " "}).SetResult(&calcexpr).Post("/calculate")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusBadRequest))

	})
	It("returns bad request status code when mismatched operations", func() {
		controller.EXPECT().Calculate("1+3--2").Return(float64(0), calculation.NewErrOperationOrder())
		resp, err := req.SetBody(api.RequestBody{Expression: "1+3--2"}).SetResult(&calcexpr).Post("/calculate")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusBadRequest))

	})
	It("returns bad request status code when floating number is invalid", func() {
		controller.EXPECT().Calculate("5..3-2").Return(float64(0), calculation.NewErrInvalidFloatingNumber())
		resp, err := req.SetBody(api.RequestBody{Expression: "5..3-2"}).SetResult(&calcexpr).Post("/calculate")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusBadRequest))

	})
	It("returns bad request status code when there is invalid symbol", func() {
		controller.EXPECT().Calculate("8*3a").Return(float64(0), calculation.NewErrInvalidSymbol())
		resp, err := req.SetBody(api.RequestBody{Expression: "8*3a"}).SetResult(&calcexpr).Post("/calculate")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusBadRequest))

	})
	It("returns bad request status code when mismatched brackets", func() {
		controller.EXPECT().Calculate("4-1)/3").Return(float64(0), calculation.NewErrMismatchedBrackets())
		resp, err := req.SetBody(api.RequestBody{Expression: "4-1)/3"}).SetResult(&calcexpr).Post("/calculate")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusBadRequest))

	})
	It("returns bad request status code when invalid value is passed to database", func() {
		resp, err := req.SetBody(`{"Expression": 1"3*2+1"}`).SetResult(&databaseResponse).Post("/create")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusBadRequest))
	})
	It("successfully push into database and return correct status code", func() {
		controller.EXPECT().Insert(postgres.Model{Expression: "1+2"}).Return(nil)
		resp, err := req.SetBody(api.RequestBody{Expression: "1+2"}).SetResult(&databaseResponse).Post("/create")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		Expect(databaseResponse.Message).To(Equal("successfully inserted row into database"))
	})
	It("returns internal server error status code when insertion in database fails", func() {
		controller.EXPECT().Insert(postgres.Model{Expression: "5-3*2"}).Return(errors.New("failed to interact with database"))
		resp, err := req.SetBody(api.RequestBody{Expression: "5-3*2"}).SetResult(&databaseResponse).Post("/create")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusInternalServerError))
	})
	It("returns correct status code and message when invalid id has been selected", func() {
		controller.EXPECT().SelectResult(int64(1)).Return(float64(0), calculator.NewInvalidIDError())
		resp, err := req.SetResult(&databaseResponse).Get("/get/1")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusBadRequest))

	})
	It("returns correct status code and message when id has not been evaluated yet", func() {
		controller.EXPECT().SelectResult(int64(1)).Return(float64(0), calculator.NewEvaluationError())
		resp, err := req.SetResult(&databaseResponse).Get("/get/1")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusBadRequest))

	})
	It("returns internal server error status code when select result from database fails", func() {
		controller.EXPECT().SelectResult(int64(1)).Return(float64(0), errors.New("failed to interact with database"))
		resp, err := req.SetResult(&databaseResponse).Get("/get/1")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusInternalServerError))
	})
	It("returns correct status code and result when result is found for id ", func() {
		controller.EXPECT().SelectResult(int64(2)).Return(float64(12), nil)
		resp, err := req.SetResult(&calcexpr).Get("/get/2")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		Expect(calcexpr.Result).To(Equal(float64(12)))
	})
	It("returns correct status code and error from evaluation when result is found for id ", func() {
		controller.EXPECT().SelectResult(int64(2)).Return(float64(0), calculator.NewCalculatorError())
		resp, err := req.SetResult(&databaseErrResponse).Get("/get/2")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusOK))
	})

})
