package modelresponses

import "net/http"

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Errors []Error     `json:"errors"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type HttpResponse struct {
	HttpStatusCode int      `json:"httpStatusCode"`
	Response       Response `json:"response"`
}

func SetHttpResponse(httpStatusCode int, data interface{}, errors []Error) HttpResponse {
	return HttpResponse{
		HttpStatusCode: httpStatusCode,
		Response: Response{
			Data:   data,
			Errors: errors,
		},
	}
}

func SetMessageHttpResponse(message string) HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusOK,
		Response: Response{
			Data: MessageResponse{
				Message: message,
			},
			Errors: []Error{},
		},
	}
}
