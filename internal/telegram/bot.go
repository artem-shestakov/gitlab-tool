package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var TgURL = "https://api.telegram.org/bot%s/"

type Bot struct {
	token       string
	apiEndpoint string
	httpClient  http.Client
}

func NewBot(token string) *Bot {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/", token)
	httpClient := http.Client{}
	return &Bot{
		token:       token,
		apiEndpoint: url,
		httpClient:  httpClient,
	}
}

func (b *Bot) MakeRequest(method string, params url.Values) (Response, error) {
	req, err := http.NewRequest("POST", b.apiEndpoint+method, strings.NewReader(params.Encode()))
	if err != nil {
		return Response{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := b.httpClient.Do(req)
	if err != nil {
		return Response{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var apiResp Response
	err = json.Unmarshal(respBody, &apiResp)
	if err != nil {
		return Response{}, err
	}
	return apiResp, nil
}

type Response struct {
	Ok     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
}

type User struct {
	ID                     int    `json:"id"`
	IsBot                  bool   `json:"is_bot"`
	FirstName              string `json:"first_name"`
	Username               string `json:"username"`
	CanJoinGroups          bool   `json:"can_join_groups"`
	CanReadAllGroupMessage bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries  bool   `json:"supports_inline_queries"`
}
