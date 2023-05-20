# opaitokens
A golang lib to help you to get openai access token and refresh the token

# How to use?
```go

email := "xxx@example.com"
password := "xxxxxxx"

tokens := NewOpaiTokens(email, password)
token := tokens.FetchToken()
fmt.Printf("token info: %v\n", token)

token = tokens.RefreshToken()
fmt.Println("token info again: %v\n", token)
	


```
