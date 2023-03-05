package helpers

import "github.com/asaskevich/govalidator"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ApiResponse(code int, status string, data interface{}, message string) Response {
	meta := Meta{
		Code:    code,
		Status:  status,
		Message: message,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(govalidator.Errors).Errors() {
		errors = append(errors, e.Error())
	}

	return errors
}
