package api

type Response struct {
	Data    any    `json:"data,omitempty"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Success(data any) Response {
	return Response{
		Data:    data,
		Code:    "200",
		Message: "success",
	}
}
