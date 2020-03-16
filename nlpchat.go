package ai

import (
	"net/url"

	"github.com/pkg/errors"
)

const (
	URLChat = "https://api.ai.qq.com/fcgi-bin/nlp/nlp_textchat"
)

type NLPChatData struct {
	Session string `json:"session"`
	Answer  string `json:"answer"`
}

type NLPChatResult struct {
	Return  int          `json:"ret"`
	Message string       `json:"msg"`
	Data    *NLPChatData `json:"data"`
}

// NLPChat 智能闲聊
func NLPChat(session string, question string) (*NLPChatData, error) {
	return app.NLPChat(session, question)
}

// NLPChat 智能闲聊
func (a *App) NLPChat(session string, question string) (*NLPChatData, error) {
	param := url.Values{}
	param.Add(KeyQuestion, question)
	param.Add(KeySession, session)
	a.prepareRequestParam(&param)
	var result NLPChatResult

	if err := a.do(URLChat, &param, &result); err != nil {
		return nil, errors.Wrap(err, "do request")
	}

	if result.Return != 0 {
		return nil, errors.Errorf("server return error=%d msg=%s", result.Return, result.Message)
	}
	return result.Data, nil
}
