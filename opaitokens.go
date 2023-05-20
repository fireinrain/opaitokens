package opaitokens

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fireinrain/opaitokens/auth"
	"github.com/fireinrain/opaitokens/model"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const OpenaiTokenBaseUrl = "https://auth0.openai.com/oauth/token"

type OpaiTokens struct {
	Email          string                     `json:"email"`
	Password       string                     `json:"password"`
	OpenaiToken    model.OpenaiToken          `json:"openaiToken"`
	RefreshedToken model.OpenaiRefreshedToken `json:"refreshedToken"`
}

func NewOpaiTokens(email string, password string) *OpaiTokens {
	if email == "" {
		log.Fatal("email cannot be empty")
	}
	if password == "" {
		log.Fatal("password cannot be empty")
	}
	return &OpaiTokens{
		Email:          email,
		Password:       password,
		OpenaiToken:    model.OpenaiToken{},
		RefreshedToken: model.OpenaiRefreshedToken{},
	}
}

func (receiver *OpaiTokens) FetchToken() *OpaiTokens {
	auth := auth.NewAuth0(receiver.Email, receiver.Password, false)
	codeAndUrl, err := auth.AuthForCodeUrl()
	if err != nil {
		fmt.Println("Error:", err)
		return receiver
	}
	codeVeriferAndCode := strings.Split(codeAndUrl, "|")
	codeVerifer := codeVeriferAndCode[0]
	code := codeVeriferAndCode[1]
	token, err := receiver.reqForToken(code, codeVerifer)
	if err == nil {
		receiver.OpenaiToken = token
	}
	return receiver
}

func (receiver *OpaiTokens) RefreshToken() *OpaiTokens {
	if receiver.OpenaiToken.RefreshToken == "" {
		receiver.FetchToken()
	} else {
		refreshTokenStr := receiver.OpenaiToken.RefreshToken
		token, err := receiver.refreshToken(refreshTokenStr)
		if err == nil {
			receiver.RefreshedToken = token
		}
	}
	return receiver
}

func (reciver *OpaiTokens) reqForToken(code string, codeVerifier string) (model.OpenaiToken, error) {
	var token model.OpenaiToken
	url := OpenaiTokenBaseUrl // 替换为实际的目标URL

	req := model.NewOpenaiTokenReq()
	req.Code = code
	req.CodeVerifier = codeVerifier

	// 构建POST请求数据
	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return token, err
	}

	resp, err := makePostRequest(url, jsonData)
	if err != nil {
		fmt.Println("makePost request error:", err)
		return token, err
	}
	err = json.Unmarshal([]byte(resp), &token)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
		return token, err
	}

	return token, nil
}

func (reciver *OpaiTokens) refreshToken(refreshToken string) (model.OpenaiRefreshedToken, error) {
	var refreshedToken model.OpenaiRefreshedToken
	url := OpenaiTokenBaseUrl // 替换为实际的目标URL

	req := model.NewOpenaiRefreshTokenReq()
	req.RefreshToken = refreshToken
	// 构建POST请求数据

	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return refreshedToken, err
	}
	resp, err := makePostRequest(url, jsonData)
	if err != nil {
		fmt.Println("error for request:", err)
		return refreshedToken, err

	}
	err = json.Unmarshal([]byte(resp), &refreshedToken)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
		return refreshedToken, err

	}
	return refreshedToken, nil
}

func makePostRequest(url string, jsonData []byte) (resp string, error error) {
	// 创建请求
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("failed to create request:", err)
		return "", err
	}
	// 设置User-Agent头部字段
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	// 发送请求
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("POST request error:", err)
		return "", err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return buf.String(), nil
}

func getLoginUrl(codeChallenge string) string {
	encodedString := "https://auth0.openai.com/authorize?client_id=pdlLIX2Y72MIl2rhLhTE9VV9bN905kBh&audience=https%3A%2F%2Fapi.openai.com%2Fv1&redirect_uri=com.openai.chat%3A%2F%2Fauth0.openai.com%2Fios%2Fcom.openai.chat%2Fcallback&scope=openid%20email%20profile%20offline_access%20model.request%20model.read%20organization.read%20offline&response_type=code&code_challenge=w6n3Ix420Xhhu-Q5-mOOEyuPZmAsJHUbBpO8Ub7xBCY&code_challenge_method=S256&prompt=login"
	//fmt.Println("decoded string:", decodedString)
	re := regexp.MustCompile(`code_challenge=[^&]+`)
	replacement := "code_challenge=" + codeChallenge
	newURL := re.ReplaceAllString(encodedString, replacement)
	//fmt.Println("Modified URL:", newURL)
	//fmt.Println(escape)
	return newURL
}
