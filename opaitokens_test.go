package opaitokens

import (
	"fmt"
	"testing"
)

func TestUserCase(t *testing.T) {
	email := "xxx@example.com"
	password := "xxxxxxx"

	tokens := NewOpaiTokens(email, password)
	token := tokens.FetchToken()
	fmt.Printf("token info: %v\n", token)

	token = tokens.RefreshToken()
	fmt.Println("token info again: %v\n", token)

}
