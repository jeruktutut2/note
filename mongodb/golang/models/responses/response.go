package modelresponses

import "net/http"

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Message struct {
	Message string `json:"message"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Errors []Error     `json:"errors"`
}

type HttpResponse struct {
	HttpStatusCode int
	Response       Response
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

func SetMessageHttpResponse(httpStatusCode int, message string) HttpResponse {
	return HttpResponse{
		HttpStatusCode: httpStatusCode,
		Response: Response{
			Data: Message{
				Message: message,
			},
			Errors: []Error{},
		},
	}
}

func SetBadRequestResponse(field string, message string) HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusBadRequest,
		Response: Response{
			Data: nil,
			Errors: []Error{
				{
					Field:   field,
					Message: message,
				},
			},
		},
	}
}

func SetNotFoundHttpResponse(field string, message string) HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusNotFound,
		Response: Response{
			Data: nil,
			Errors: []Error{
				{
					Field:   field,
					Message: message,
				},
			},
		},
	}
}

func SetUnauthorizedHttpResponse(field string, message string) HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusUnauthorized,
		Response: Response{
			Data: nil,
			Errors: []Error{
				{
					Field:   field,
					Message: message,
				},
			},
		},
	}
}

func SetUserCloseHttpConnectionErrorResponse() HttpResponse {
	return HttpResponse{
		HttpStatusCode: 499,
		Response: Response{
			Data: nil,
			Errors: []Error{
				{
					Field:   "message",
					Message: "user close http connection or cancel http connection",
				},
			},
		},
	}
}

func SetTimeoutErrorResponse() HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusRequestTimeout,
		Response: Response{
			Data: nil,
			Errors: []Error{
				{
					Field:   "message",
					Message: "request timeout",
				},
			},
		},
	}
}

func SetRefreshTokenExpiredHttpResponse() HttpResponse {
	return HttpResponse{
		HttpStatusCode: 498,
		Response: Response{
			Data: nil,
			Errors: []Error{
				{
					Field:   "message",
					Message: "refresh token has expired",
				},
			},
		},
	}
}

func SetInternalServerErrorResponse() HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusInternalServerError,
		Response: Response{
			Data: nil,
			Errors: []Error{
				{
					Field:   "message",
					Message: "internal server error",
				},
			},
		},
	}
}
