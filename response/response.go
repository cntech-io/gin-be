package response

import "net/http"

type responseMessage string

const (
	SUCCESS responseMessage = "SUCCESS"
	FAIL    responseMessage = "FAIL"
)

type response struct {
	Status         bool            `json:"status"`
	Message        responseMessage `json:"message"`
	EndUserMessage string          `json:"end_user_message"`
	Error          string          `json:"error"`
	Data           any             `json:"data,omitempty"`
}

func Response(endUserMessage string) *response {
	return &response{
		EndUserMessage: endUserMessage,
	}
}

func (r *response) Success() *response {
	r.Status = true
	r.Message = SUCCESS
	return r
}

func (r *response) Fail(err string) *response {
	r.Status = true
	r.Message = SUCCESS
	r.Error = err
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

func ForbittenResponse(endUserMessage string, err string) (int, response) {
	return http.StatusForbidden, *Response(endUserMessage).Fail(err)
}

func BadRequestResponse(endUserMessage string, err string) (int, response) {
	return http.StatusBadRequest, *Response(endUserMessage).Fail(err)
}

func UnauthorizedResponse(endUserMessage string, err string) (int, response) {
	return http.StatusUnauthorized, *Response(endUserMessage).Fail(err)
}
