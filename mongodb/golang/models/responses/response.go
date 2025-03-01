package modelresponses

import "net/http"

type MessageResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

type HttpResponse struct {
	HttpStatusCode int
	Response       Response
}

func SetHttpResponse(httpStatusCode int, data interface{}, errors interface{}) HttpResponse {
	return HttpResponse{
		HttpStatusCode: httpStatusCode,
		Response: Response{
			Data:   data,
			Errors: errors,
		},
	}
}

func SetDataHttpResponse(httpStatusCode int, data interface{}) HttpResponse {
	return HttpResponse{
		HttpStatusCode: httpStatusCode,
		Response: Response{
			Data:   data,
			Errors: nil,
		},
	}
}

func SetMessageHttpResponse(httpStatusCode int, message string) HttpResponse {
	return HttpResponse{
		HttpStatusCode: httpStatusCode,
		Response: Response{
			Data: MessageResponse{
				Message: message,
			},
			Errors: nil,
		},
	}
}

func SetBadRequestResponse(message string) HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusBadRequest,
		Response: Response{
			Data: nil,
			Errors: MessageResponse{
				Message: message,
			},
		},
	}
}

func SetNotFoundHttpResponse(message string) HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusNotFound,
		Response: Response{
			Data: nil,
			Errors: MessageResponse{
				Message: message,
			},
		},
	}
}

func SetUnauthorizedHttpResponse(message string) HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusUnauthorized,
		Response: Response{
			Data: nil,
			Errors: MessageResponse{
				Message: message,
			},
		},
	}
}

func SetUserCloseHttpConnectionErrorResponse() HttpResponse {
	return HttpResponse{
		HttpStatusCode: 499,
		Response: Response{
			Data: nil,
			Errors: MessageResponse{
				Message: "user close http connection or cancel http connection",
			},
		},
	}
}

func SetTimeoutErrorResponse() HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusRequestTimeout,
		Response: Response{
			Data: nil,
			Errors: MessageResponse{
				Message: "request timeout",
			},
		},
	}
}

func SetRefreshTokenExpiredHttpResponse() HttpResponse {
	return HttpResponse{
		HttpStatusCode: 498,
		Response: Response{
			Data: nil,
			Errors: MessageResponse{
				Message: "refresh token has expired",
			},
		},
	}
}

func SetInternalServerErrorResponse() HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusInternalServerError,
		Response: Response{
			Data: nil,
			Errors: MessageResponse{
				Message: "internal server error",
			},
		},
	}
}
