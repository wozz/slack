package slack

import (
	"encoding/json"
	"errors"
)

type Emoji struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

// API emoji.list: Lists all custom emojis in a Slack team.
func (sl *Slack) EmojisList() ([]*Emoji, error) {
	uv := sl.urlValues()
	body, err := sl.GetRequest(emojiListApiEndpoint, uv)
	if err != nil {
		return nil, err
	}
	res := new(EmojiListAPIResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.Error)
	}
	return res.Emojis()
}

// response type of `users.list` api
type EmojiListAPIResponse struct {
	BaseAPIResponse
	RawEmojis json.RawMessage `json:"emoji"`
}

func (res *EmojiListAPIResponse) Emojis() ([]*Emoji, error) {
	var emojis []*Emoji
	err := json.Unmarshal(res.RawEmojis, &emojis)
	if err != nil {
		return nil, err
	}
	return emojis, nil
}

