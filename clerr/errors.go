package clerr

import "github.com/pkg/errors"

// Auth error
const (
	errorServer = "server error"
)

var ErrorServer = errors.New(errorServer)
