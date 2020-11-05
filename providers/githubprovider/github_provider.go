package githubprovider

import (
	"fmt"
)

const (
	token                     = ""
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
)

// GetAuthorizationHeader returns the Authorization header for authenticating to Github
func GetAuthorizationHeader(token string) string {
	// Header {"Authorization": "token 0000000000000000000000000000000000000000"}
	return fmt.Sprintf(headerAuthorizationFormat, token)
}
