# opaitokens
A golang lib to help you to get openai access token and refresh the token

# How to use?
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
