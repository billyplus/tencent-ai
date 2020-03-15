package ai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	ChatURL      = "https://api.ai.qq.com/fcgi-bin/nlp/nlp_textchat"
	KeyAppID     = "app_id"
	KeyTimeStamp = "time_stamp"
	KeyNonceStr  = "nonce_str"
	KeySign      = "sign"
	KeySession   = "session"
	KeyQuestion  = "question"
)

type NLPData struct {
	Session string `json:"session"`
	Answer  string `json:"answer"`
}

type NLPResult struct {
	Return  int      `json:"ret"`
	Message string   `json:"msg"`
	Data    *NLPData `json:"data"`
}

func NLPChat(session string, question string) (*NLPData, error) {
	return app.NLPChat(session, question)
}

func (a *App) NLPChat(session string, question string) (*NLPData, error) {
	param := url.Values{}
	param.Add(KeyAppID, a.idStr)
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	param.Add(KeyTimeStamp, timestamp)
	param.Add(KeyNonceStr, "teststr")
	param.Add(KeySession, session)
	param.Add(KeyQuestion, question)
	// param.Add(KeySign, "")
	sign := GetRequestSign(param, a.key)
	param.Set(KeySign, sign)

	resp, err := a.client.Post(ChatURL, "application/x-www-form-urlencoded", strings.NewReader(param.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "post request")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("response code=%d status=%s", resp.StatusCode, resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}
	var result NLPResult
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, errors.Wrap(err, "Unmarshal result")
	}
	if result.Return != 0 {
		return nil, errors.Errorf("server return error=%d msg=%s", result.Return, result.Message)
	}
	return result.Data, nil
}
