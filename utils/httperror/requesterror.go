package httperror

import "net/http"

type RestErr struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys"`
	Message     string `json:"message"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Code:        http.StatusBadRequest,
		MessageKeys: "invalid body",
		Message:     message,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message:     message,
		Code:        http.StatusInternalServerError,
		MessageKeys: "Internal server error",
	}
}

func NewUnAuthenticatedServerError(message string) *RestErr {
	return &RestErr{
		Message:     message,
		Code:        http.StatusUnauthorized,
		MessageKeys: "user_is_unauthorized",
	}
}
