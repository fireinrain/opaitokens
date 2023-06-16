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
```


## use pooled token with ai.fakenopen.com

```go
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

```

## renew shared token for keep pooled token valid
```go
//主动在14天之内刷新所有账号的shared token 来确保pooled token有效
//可以使用
var accounts []OpenaiAccount
account := OpenaiAccount{
Email:    "xxxx@gmail.com",
Password: "xx@xx",
MFA:      "",
}
accounts = append(accounts, account)
tokens := FakeOpenTokens{}
renewResult, err := tokens.RenewSharedToken(accounts)
if err != nil {
    fmt.Println("error: ", err)
}
fmt.Println(renewResult)

```

