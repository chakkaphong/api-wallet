package responses

import "api-wallet/internal/global/errors"

var ErrorCodeMapping = map[int]string{
	400: "bad_request",
	401: "unauthorized",
	403: "forbidden",
	404: "data_not_found",
	500: "internal_server_error",
}

func BadRequest(msg ...string) (res Responses) {
	m := "Bad request."
	if len(msg) > 0 {
		m = msg[0]
	}
	return Responses{
		Code:    "bad_request",
		Message: m,
		Data:    nil,
	}
}

func ResponseFailed(err error) (res Responses) {
	appErr, ok := err.(errors.AppError)
	if ok {
		codeStr := ErrorCodeMapping[appErr.Code]

		return Responses{
			Code:    codeStr,
			Message: appErr.Message,
			Data:    nil,
		}
	}

	return Responses{
		Code:    "internal_server_error",
		Message: "internal_server_error",
		Data:    nil,
	}
}
