package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// GenerateCodeVerifier
//
//	@Description: 生成校验code
//	@return string
func GenerateCodeVerifier() string {
	// 随机生成一个长度为 32 的 code_verifier
	token := make([]byte, 32)
	rand.Read(token)
	codeVerifier := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(token)
	return codeVerifier
}

// GenerateCodeChallenge
//
//	@Description: 校验code challenge
//	@param codeVerifier
//	@return string
func GenerateCodeChallenge(codeVerifier string) string {
	// 对 code_verifier 进行哈希处理，然后再进行 base64url 编码，生成 code_challenge
	sha256Hash := sha256.Sum256([]byte(codeVerifier))
	codeChallenge := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sha256Hash[:])
	return codeChallenge
}
