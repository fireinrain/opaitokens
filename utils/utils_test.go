package utils

import (
	"fmt"
	"testing"
)

func TestGenerateCodeVerifier(t *testing.T) {
	verifier := GenerateCodeVerifier()
	fmt.Println(verifier)
}

func TestGenerateCodeChallenge(t *testing.T) {
	verifier := GenerateCodeVerifier()
	challenge := GenerateCodeChallenge(verifier)
	fmt.Println(challenge)
}
