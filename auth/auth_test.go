package auth

import (
	"fmt"
	"testing"
)

func TestAuthForToken(t *testing.T) {
	email := "xxx@example.com"
	password := "xxxx"
	useCache := false

	auth := NewAuth0(email, password, useCache)
	accessToken, err := auth.Auth(true)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Access Token:", accessToken)
}

func TestAuthForTokenByProxy(t *testing.T) {
	email := "xxx@example.com"
	password := "xxxx"
	useCache := false

	auth := NewAuth0(email, password, useCache)
	accessToken, err := auth.Auth(false)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Access Token:", accessToken)
}
