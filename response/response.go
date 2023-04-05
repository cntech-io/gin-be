package response

import "net/http"

type responseType string

const (
	InvalidParams  responseType = "invalid_params"
	QueryError     responseType = "query_error"
	AlgorthymError responseType = "algorthym_error"
)

type response struct {
	Status         bool   `json:"status"`
	EndUserMessage string `json:"end_user_message"`
	Type           string `json:"type"`
	Data           any    `json:"data,omitempty"`
}

func Response(endUserMessage string) *response {
	return &response{
		EndUserMessage: endUserMessage,
	}
}

func (r *response) Success() *response {
	r.Status = true
	return r
}

func (r *response) Fail(respType responseType) *response {
	r.Status = false
	r.Type = string(respType)
	return r
}
func (r *response) WithData(data any) *response {
	r.Data = data
	return r
}

func OKResponse(endUserMessage string) (int, response) {
	return http.StatusOK, *Response(endUserMessage).Success()
}

func DataResponse(endUserMessage string, data any) (int, response) {
	return http.StatusOK, *Response(endUserMessage).Success().WithData(data)
}

func CreatedResponse(endUserMessage string) (int, response) {
	return http.StatusCreated, *Response(endUserMessage).Success()
}

func ForbittenResponse(endUserMessage string, respType responseType) (int, response) {
	return http.StatusForbidden, *Response(endUserMessage).Fail(respType)
}

func BadRequestResponse(endUserMessage string, respType responseType) (int, response) {
	return http.StatusBadRequest, *Response(endUserMessage).Fail(respType)
}

func UnauthorizedResponse(endUserMessage string, respType responseType) (int, response) {
	return http.StatusUnauthorized, *Response(endUserMessage).Fail(respType)
}
