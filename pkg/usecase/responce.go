package usecase

type FailedResponseBody struct {
	ErrorCode ErrorCode   `json:"errorCode"`
	Message   interface{} `json:"message"`
}

type FailedResponse struct {
	HttpCode int16
	Message  interface{}
}

var responseGroups = map[int16][]ErrorCode{
	200: {Success},
	201: {ResourceCreated},
	204: {NoContent},
	400: {BadRequest},
	401: {Unauthorized, UnregisteredAccount},
	403: {Forbidden, InvalidAccountStatus},
	404: {NotFound},
	409: {Conflict},
	432: {BadHeader},
	433: {ResourceAlreadyExist},
	434: {ResourceInTrash},
	500: {InternalServerError},
}

var ErrorCodeToFailedResponse = ConvertResponseGroups(responseGroups)

func ConvertResponseGroups(groups map[int16][]ErrorCode) map[ErrorCode]FailedResponse {
	var result = make(map[ErrorCode]FailedResponse)
	for key, errorCodes := range groups {
		for _, code := range errorCodes {
			result[code] = FailedResponse{
				HttpCode: key,
				Message:  code.Message(),
			}
		}
	}
	return result
}

func (e ErrorCode) Message() interface{} {
	return [...]interface{}{
		"success",
		"create",
		"resourceDeleted",
		EmptyResponse,
		"badRequest",
		"unauthorized",
		"forbidden",
		"badHeader",
		"conflict",
		"resourceAlreadyExist",
		"internalServerError",
		"resourceDeleted",
		"accountIsNotRegistered",
		"invalidAccountStatus",
		"confirmationCodeIsTimeOut",
		"confirmationCodeIsNotMatched",
		"accountMustBeConfirmedBefore",
		"notFound",
		"resourceInTrash",
	}[e]
}
