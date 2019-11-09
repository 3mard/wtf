package pocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/wtfutil/wtf/logger"
)

/*
	Authorization workflow is documented at https://getpocket.com/developer/docs/authentication
	broken to 4 steps :
		1- Obtain a platform consumer key from http://getpocket.com/developer/apps/new.
		2- Obtain a request token
		3- Redirect user to Pocket to continue authorization
		4- Receive the callback from Pocket, this wont be used
		5- Convert a request token into a Pocket access token
*/

// Client is pocket api client Documention at https://getpocket.com/developer/docs/overview
type Client struct {
	consumerKey string
	baseUr      string
	redirectUrl string
}

func NewClient(consumerKey, redirectUrl string) *Client {
	return &Client{
		consumerKey: consumerKey,
		redirectUrl: redirectUrl,
		baseUr:      "https://getpocket.com/v3",
	}

}

type ObtainRequestTokenRequest struct {
	ConsumerKey string `json:"consumer_key"`
	RedirectURI string `json:"redirect_uri"`
}
type AccessTokenRequest struct {
	ConsumerKey string `json:"consumer_key"`
	RequestCode string `json:"code"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token",omitempty`
}

func (client *Client) request(url string, requestBody interface{}, result interface{}) error {

	jsonValues, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValues))
	if err != nil {
		return err
	}

	request.Header.Add("X-Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	logger.Log(string(responseBody))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseBody, &result); err != nil {
		return fmt.Errorf(`Could not unmarshal url [%s] 
		response [%s] request[%s] error:%v`, url, responseBody, jsonValues, err)
	}

	return nil

}

func (client *Client) ObtainRequestToken() (code string, err error) {
	url := fmt.Sprintf("%s/oauth/request", client.baseUr)
	requestData := ObtainRequestTokenRequest{ConsumerKey: client.consumerKey, RedirectURI: client.redirectUrl}

	var responseData map[string]string

	err = client.request(url, requestData, &responseData)

	if err != nil {
		return code, err
	}

	return responseData["code"], nil

}

func (client *Client) CreateRedirectLink(requestToken string) string {
	return fmt.Sprintf("https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s", requestToken, client.redirectUrl)
}

func (client *Client) getAccessToken(requestToken string) (accessToken string, err error) {
	url := fmt.Sprintf("%s/oauth/authorize", client.baseUr)
	requestData := AccessTokenRequest{
		ConsumerKey: client.consumerKey,
		RequestCode: requestToken,
	}

	var response AccessTokenResponse
	err = client.request(url, requestData, &response)
	if err != nil {
		return "", err
	}
	return response.AccessToken, nil

}
