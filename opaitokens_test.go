package opaitokens

import (
	"fmt"
	"testing"
)

func TestUserCase(t *testing.T) {
	email := "xxxx@xx.com"
	password := "xxxxx"

	tokens := NewOpaiTokens(email, password, true)
	token := tokens.FetchToken()
	fmt.Printf("token info: %v\n", token)
	accessToken := token.OpenaiToken.AccessToken
	// use the access token
	fmt.Printf("i am using access token: %v \n", accessToken)

	token = tokens.RefreshToken()
	fmt.Printf("token info again: %v\n", token)
	accessToken = token.RefreshedToken.AccessToken
	//use the refresh token
	fmt.Println("i am using refresh token: ", accessToken)

}

func TestNewOpaiTokensWithMFA(t *testing.T) {
	email := "xxxx@xx.com"
	password := "xxxxx"
	mfa := "your mfa code"

	tokens := NewOpaiTokensWithMFA(email, password, mfa, false)
	token := tokens.FetchToken()
	fmt.Printf("token info: %v\n", token)
	accessToken := token.OpenaiToken.AccessToken
	// use the access token
	fmt.Printf("i am using access token: %v \n", accessToken)

	token = tokens.RefreshToken()
	fmt.Printf("token info again: %v\n", token)
	accessToken = token.RefreshedToken.AccessToken
	//use the refresh token
	fmt.Println("i am using refresh token: ", accessToken)
}

func TestFakeOpenTokens_FetchSharedToken(t *testing.T) {
	tokens := FakeOpenTokens{}
	account := OpenaiAccount{
		Email:    "xxxx@gmail.com",
		Password: "xx@xx",
		MFA:      "",
	}
	token, err := tokens.FetchSharedToken(account, "fireinrain")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(token.TokenKey)
}

func TestFakeOpenTokens_FetchPooledToken(t *testing.T) {
	var accounts []OpenaiAccount
	account := OpenaiAccount{
		Email:    "xxxx@gmail.com",
		Password: "xx@xx",
		MFA:      "",
	}
	accounts = append(accounts, account)
	tokens := FakeOpenTokens{}
	token, err := tokens.FetchPooledToken(accounts)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(token)
}

func TestFakeOpenTokens_FetchMixedPooledToken(t *testing.T) {
	var accounts []OpenaiAccount
	account := OpenaiAccount{
		Email:    "xxxx@gmail.com",
		Password: "xx@xx",
		MFA:      "",
	}
	accounts = append(accounts, account)

	var skKeys []string
	skKeys = append(skKeys, "sk-xxxxxx")
	tokens := FakeOpenTokens{}
	token, err := tokens.FetchMixedPooledToken(accounts, skKeys)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(token)
}
