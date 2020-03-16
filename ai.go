package ai

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type App struct {
	id     int
	idStr  string
	key    string
	client *http.Client
}

var (
	app = NewApp(0, "")
)

const (
	KeyAppID     = "app_id"
	KeyTimeStamp = "time_stamp"
	KeyNonceStr  = "nonce_str"
	KeySign      = "sign"
	KeySession   = "session"
	KeyQuestion  = "question"
	KeyText      = "text"
	KeySource    = "source"
	KeyTarget    = "target"
)

func NewApp(appid int, key string) *App {
	app := &App{
		id:     appid,
		key:    key,
		idStr:  fmt.Sprintf("%d", appid),
		client: &http.Client{},
	}
	return app
}

func ResetDefaultApp(appid int, key string) {
	app.id = appid
	app.idStr = fmt.Sprintf("%d", appid)
	app.key = key
}

func randomStr() string {
	v := uint64(rand.Int63())
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, v)
	return hex.EncodeToString(b)
}

func (a *App) prepareRequestParam(param *url.Values) {
	param.Add(KeyAppID, a.idStr)
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	param.Add(KeyTimeStamp, timestamp)
	param.Add(KeyNonceStr, randomStr())
	// param.Add("sign", "")
	sign := GetRequestSign(param, a.key)
	param.Add(KeySign, sign)
	fmt.Println(param.Encode())
}

func bindJSON(reader io.ReadCloser, value interface{}) error {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}
	if err = json.Unmarshal(body, value); err != nil {
		return errors.Wrap(err, "Unmarshal result")
	}
	return nil
}

func (a *App) do(url string, param *url.Values, result interface{}) error {
	resp, err := a.client.Post(url, "application/x-www-form-urlencoded", strings.NewReader(param.Encode()))
	if err != nil {
		return errors.Wrap(err, "post request")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("response code=%d status=%s", resp.StatusCode, resp.Status)
	}
	if err = bindJSON(resp.Body, &result); err != nil {
		return errors.Wrap(err, "failed to unmarshal response")
	}
	return nil
}
