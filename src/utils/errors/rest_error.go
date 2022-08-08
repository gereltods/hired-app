package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// type RestErr interface {
// 	Message() string
// 	Status() int
// 	Error() string
// 	Causes() []interface{}
// }

type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

// func (e restErr) Error() string {
// 	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
// 		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
// }

// func (e restErr) Message() string {
// 	return e.ErrMessage
// }

// func (e restErr) Status() int {
// 	return e.ErrStatus
// }

// func (e restErr) Causes() []interface{} {
// 	return e.ErrCauses
// }

func NewError(msg string) error {
	return errors.New(msg)
}

func NewRestError(message string, status int, err string, causes []interface{}) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Error:   err,
		Causes:  causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr RestErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return apiErr, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
