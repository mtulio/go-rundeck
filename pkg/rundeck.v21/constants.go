package rundeck

import "errors"

// MaxRundeckVersion is the maximum version of the api this library supports
// can be overridden
const MaxRundeckVersion = "21"

const (
	basicAuthType = "basic"
)

var (
	// ErrInvalidUsernamePassword is the error typed returned when basic auth fails
	ErrInvalidUsernamePassword = errors.New("Invalid username or password returned by rundeck")

	// ErrInvalidRundeckURL is the error for an invalid rundeck url
	ErrInvalidRundeckURL = errors.New("Invalid Rundeck URL")

	// ErrAuthFailed is the error for a auth failure in an api call
	// this is slightly different the ErrInvalidUsernamePassword
	// as this means auth succeeded with basic auth but a 401 could be returned farther down
	ErrAuthFailed = errors.New("API call failed due to authentication failure")

	// ErrMissingResource is the error type for 404 not found
	ErrMissingResource = errors.New("Rundeck could not find the resource you requested")
)