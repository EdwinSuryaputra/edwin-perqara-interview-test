package response

type HandlerResponse[T any] struct {
	Data T `json:"data"`
}

type HandlerErrorResponse struct {
	Error HandlerErrorResponseData `json:"error"`
}

type HandlerErrorResponseData struct {
	Message string `json:"message"`
}

func HandlerRespondJsonSuccess[T any](data T) HandlerResponse[T] {
	return HandlerResponse[T]{
		Data: data,
	}
}

func HandlerResponseJsonError(message string) HandlerErrorResponse {
	return HandlerErrorResponse{
		Error: HandlerErrorResponseData{
			Message: message,
		},
	}
}
