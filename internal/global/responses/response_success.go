package responses

type Responses struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(i interface{}) (res Responses) {
	return Responses{
		Code:    "success",
		Message: "success",
		Data:    i,
	}
}
