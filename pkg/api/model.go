package api

type RequestBody struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result"`
}

type DatabaseResponse struct {
	Message string `json:message`
}

type DatabaseErrorResponse struct {
	Error string `json:error`
}
