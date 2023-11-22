package models

type messageResponse struct {
	Message string `json:"message"`
}

func Err_response(err error) messageResponse {
	return messageResponse{
		Message: err.Error(),
	}
}
