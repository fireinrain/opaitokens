# opaitokens
A golang lib to help you to get openai access token and refresh the token

# How to use?

## if you cant visit openai freely, then you need set HTTP_PROXY and HTTPS_PROXY env before use.
```bash
unix/linux/macos

export http_proxy=http://proxy.example.com:port
export https_proxy=http://proxy.example.com:port

windows

set http_proxy=http://proxy.example.com:port
set https_proxy=http://proxy.example.com:port


```

## official account
```go
email := "xxxx@xx.com"
password := "xxxxx"
//false mean you dont want to use fakeopenproxy to login
// you can set it as true, so when you dont have a vps for get your access with fakeopen api proxy
tokens := NewOpaiTokens(email, password, false)
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

tokens := NewOpaiTokensWithMFA(email, password,mfa,false)
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

## fetch pooled token with official accounts and offline sk keys
```go
var accounts []OpenaiAccount
account := OpenaiAccount{
    Email:    "xxxx@gmail.com",
    Password: "xx@xx",
    MFA:      "",
}
accounts = append(accounts, account)
	
var skKeys []string
skKeys = append(skKeys,"sk-xxxxxx")
tokens := FakeOpenTokens{}
token, err := tokens.FetchMixedPooledToken(accounts,skKeys)
if err != nil {
    fmt.Println("error: ", err)
}
fmt.Println(token)


```

## if you have store the openai refresh token, congratulation! you can use openai official refresh token to get access token and then use in fakeopen fk and pk.
```go
func TestFakeOpenTokens_FetchSharedTokenWithRefreshToken(t *testing.T) {
	openaiAccountEmail := "xxxx@gmail.com"
	openaiRefreshToken := "xxxxxxxx"
	uniqueName := "fireinrain"
	receiver := FakeOpenTokens{}
	token, err := receiver.FetchSharedTokenWithRefreshToken(openaiAccountEmail, openaiRefreshToken, uniqueName)
	if err != nil{
		fmt.Println("error: ", err)
	}
	fmt.Printf("%v\n", token)

}


```
Note that every function with suffic "WithRefreshToken" will use openai official refresh token to get access token and then use in fakeopen fk and pk.

## Thanks for Cloudflare Platform.
![cloudflare](./cf.svg)

You can use cloudflare free cdn to speed your website. And if necessary, choose the pro subscription.
[Cloudflare](https://www.cloudflare.com/)

