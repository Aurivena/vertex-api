package usecase

const (
	Success ErrorCode = iota
	Create
	ResourceCreated
	NoContent
	BadRequest
	Unauthorized
	Forbidden
	BadHeader
	Conflict
	ResourceAlreadyExist
	InternalServerError
	ResourceDeleted
	UnregisteredAccount
	InvalidAccountStatus
	ConfirmationTimeout
	ConfirmCodeNotMatched
	AccountMustConfirmed
	NotFound
	ResourceInTrash
)

func (e ErrorCode) String() string {
	return [...]string{
		"success",
		"create",
		"resourceCreated",
		"noContent",
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

type ErrorCode int

var EmptyResponse interface{}
