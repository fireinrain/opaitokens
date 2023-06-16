package fakeopen

import (
	"fmt"
	"testing"
)

func TestShareToken(t *testing.T) {

	platform := AiFakeOpenPlatform{}
	req := SharedTokenReq{
		UniqueName:        "abcdxxxxxx",
		AccessToken:       "eyJhbGciOiJSUzI1NiIsFpLmNvbS9hdXRoIjp7InVzZXJwczovL2F1dGgwLm9wZW5haS5jb20vIiwic3ViIjoiZ29vZ2xlLW9hdXRoMnwxMDk5NjQwMDU4NDA3MDM1MDI3MDAiLCJhdWQiOlsiaHR0cHM6Ly9hcGkub3BlbmFpLmNvbS92MSIsImh0dHBzOi8vb3BlbmFpLm9wZW5haS5hdXRoMGFwcC5jb20vdXNlcm...",
		ExpiresIn:         0,
		SiteLimit:         "",
		ShowConversations: false,
	}
	token, err := platform.GetSharedToken(req)
	if err != nil {
		fmt.Printf("error getting shared token: %v\n", err)
	}
	fmt.Println("Shared token: ", token)

}

func TestPooledToken(t *testing.T) {
	platform := AiFakeOpenPlatform{}
	//tokens with shared token
	fk := []string{"fk-_r_ugQX_7Fe"}
	req := PooledTokenReq{
		ShareTokens: fk,
		PoolToken:   "",
	}
	token, err := platform.RenewPooledToken(req)
	if err != nil {
		fmt.Printf("error renewing pool token: %v\n", err)
	}
	fmt.Println("Pooled token: ", token)
}
