package telegram

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (b *Bot) GetMe() User {
	user := User{}
	resp, err := b.MakeRequest("getMe", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(resp.Result, &user)
	return user
}

func (b *Bot) SendMesage(chatId string, text string) {
	params := url.Values{}
	params.Set("chat_id", chatId)
	params.Set("text", text)
	resp, err := b.MakeRequest("sendMessage", params)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(resp.Ok)
}
