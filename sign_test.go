package ai

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestSign(t *testing.T) {
	values := url.Values{}
	values.Add("app_id", "10000")
	values.Add("time_stamp", "323424234")
	values.Add("nonce_str", "teststr")
	values.Add("session", "10000")
	values.Add("question", "在吗？")
	sig := GetRequestSign(&values, "dfasfdawerfdafsewr")
	assert.Equal(t, "F2CA996B5D10200B6DCE144DF2B4D452", sig, "Sign result")
}
