package token

import (
	"time"
)

// Maker is an interface for managing tokens
type Maker interface {
	// this function creates a new token for a specfic username and duration
	// returns a signed token string or an error
	// this method would create and sign a token based on a specific username
	// and signed duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// function to verify token is valid or not
	// if valid the method would return the input data stored within
	// the body of the payload
	VerifyToken(token string) (*Payload, error)
}
