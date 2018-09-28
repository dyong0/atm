package response

import "net/http"

func BadRequest() (int, errResponse) {
	return http.StatusBadRequest, Error().Withmessage("bad request")
}

func InternalServerError() (int, errResponse) {
	return http.StatusInternalServerError, Error().Withmessage("internal server error")
}

func Error() errResponse {
	return errResponse{Code: 500, Message: ""}
}

func (er errResponse) WithCode(code int) errResponse {
	er.Code = code
	return er
}

func (er errResponse) Withmessage(msg string) errResponse {
	er.Message = msg
	return er
}

type errResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
