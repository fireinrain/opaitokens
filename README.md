# opaitokens
A golang lib to help you to get openai access token and refresh the token

# How to use?

## official account
```go
email := "xxxx@xx.com"
password := "xxxxx"

tokens := NewOpaiTokens(email, password)
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

```

## official account with MFA

```go
// if you have set MFA to the account, then use this below
email := "xxxx@xx.com"
password := "xxxxx"
mfa := "your mfa code"

tokens := NewOpaiTokensWithMFA(email, password,mfa)
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

```
## use share token with ai.fakenopen.com 
```go


platform := AiFakeOpenPlatform{}
req := SharedTokenReq{
    UniqueName:        "abcdxxxxxx",
    AccessToken:       "",
    ExpiresIn:         0,
    SiteLimit:         "",
    ShowConversations: false,
}
token, err := platform.GetSharedToken(req)
if err != nil {
    fmt.Printf("error getting shared token: %v\n", err)
}
fmt.Println("Shared token: ",token)
```


## use pooled token with ai.fakenopen.com

```go
platform := AiFakeOpenPlatform{}
//tokens with shared token
fk := []string{"abc"}
req := PooledTokenReq{
    ShareTokens: fk,
    PoolToken:   "",
}
token, err := platform.RenewPooledToken(req)
if err != nil {
    fmt.Printf("error renewing pool token: %v\n",err)
}
fmt.Println("Pooled token: ",token)


```

