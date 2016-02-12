package slack

import (
	"encoding/json"
	"errors"
)

// API emoji.list: Lists all custom emojis in a Slack team.
func (sl *Slack) EmojisList() ([]*interface{}, error) {
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

type EmojiListAPIResponse struct {
	BaseAPIResponse
	RawEmojis json.RawMessage `json:"emoji"`
}

func (res *EmojiListAPIResponse) Emojis() ([]*interface{}, error) {
	var emojis []*interface{}
	err := json.Unmarshal(res.RawEmojis, &emojis)
	if err != nil {
		return nil, err
	}
	return emojis, nil
}

