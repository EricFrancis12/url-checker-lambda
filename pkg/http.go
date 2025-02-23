package pkg

import "fmt"

func BearerHeader(authToken string) string {
	return fmt.Sprintf("Bearer %s", authToken)
}
