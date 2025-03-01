package modelresponses

import "net/http"

type Response struct {
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type HttpResponse struct {
	HttpStatusCode int      `json:"httpStatusCode"`
	Response       Response `json:"response"`
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

func SetMessageHttpResponse(message string) HttpResponse {
	return HttpResponse{
		HttpStatusCode: http.StatusOK,
		Response: Response{
			Data: MessageResponse{
				Message: message,
			},
			Errors: nil,
		},
	}
}

func SetBadRequestHttpResponse(message string) HttpResponse {
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

func SetInternalServerErrorHttpResponse() HttpResponse {
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
