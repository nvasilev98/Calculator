package api

import (
	"context"
	"net/http"

	"github.com/nvasilev98/calculator/pkg/api"
	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
)

type Credentials struct {
	Username string
	Password string
}

type Response struct {
	Result float64
}

func NewCalculatorClient(client *http.Client, url string, credentials Credentials) *CalculatorClient {
	calculatorClient := resty.NewWithClient(client).SetHostURL(url).SetBasicAuth(credentials.Username, credentials.Password)

	return &CalculatorClient{
		client: calculatorClient,
	}
}

type CalculatorClient struct {
	client *resty.Client
}

func (c *CalculatorClient) GetResult(ctx context.Context, expression string) (float64, error) {
	calcexpr := api.Response{}
	resp, err := c.client.R().SetBody(api.RequestBody{Expression: expression}).SetContext(ctx).SetResult(&calcexpr).Post("/calculate")
	if err != nil {
		return 0, errors.Wrap(err, "failed to get result")
	}
	switch resp.StatusCode() {
	case http.StatusOK:
		return calcexpr.Result, nil
	case http.StatusUnauthorized:
		return 0, api.NewUnauthorizedError()
	default:
		return 0, api.NewAPIError(resp.StatusCode())
	}
}
