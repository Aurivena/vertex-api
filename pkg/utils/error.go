package utils

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

type ErrorCode int
