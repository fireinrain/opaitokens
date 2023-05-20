package model

type OpenaiToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}
type OpenaiTokenRequest struct {
	RedirectURI  string `json:"redirect_uri"`
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	Code         string `json:"code"`
	CodeVerifier string `json:"code_verifier"`
}

func NewOpenaiTokenReq() *OpenaiTokenRequest {
	return &OpenaiTokenRequest{
		RedirectURI:  "com.openai.chat://auth0.openai.com/ios/com.openai.chat/callback",
		GrantType:    "authorization_code",
		ClientID:     "pdlLIX2Y72MIl2rhLhTE9VV9bN905kBh",
		Code:         "",
		CodeVerifier: "",
	}
}

////////////////////////////OpenaiToken end////////////////////////////////////

type OpenaiTokenRereshReq struct {
	RedirectURI  string `json:"redirect_uri"`
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	RefreshToken string `json:"refresh_token"`
}

func NewOpenaiRefreshTokenReq() *OpenaiTokenRereshReq {
	return &OpenaiTokenRereshReq{
		RedirectURI:  "com.openai.chat://auth0.openai.com/ios/com.openai.chat/callback",
		GrantType:    "refresh_token",
		ClientID:     "pdlLIX2Y72MIl2rhLhTE9VV9bN905kBh",
		RefreshToken: "",
	}
}

type OpenaiRefreshedToken struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

///////////////////////////OpenaiRefreshedToken end/////////////////////////////////////
