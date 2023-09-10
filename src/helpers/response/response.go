package response

import "net/http"

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

type NotFoundResponse struct {
	Response Response `json:"error"`
}

type BadRequestResponse struct {
	Response Response `json:"error"`
}

func NewNotFoundResponse(message string) NotFoundResponse {
	return NotFoundResponse{
		Response: Response{
			Message: message,
			StatusCode: http.StatusNotFound,
		},
	}
}

func NewBadRequestResponse(message string) BadRequestResponse {
	return BadRequestResponse{
		Response: Response{
			Message: message,
			StatusCode: http.StatusBadRequest,
		},
	}
}
