package ai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNLPChat(t *testing.T) {
	ResetDefaultApp(111111, "xxxxxxxxxxx")
	data, err := app.NLPChat("100000", "你叫什么名字？？？？？？")
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, "100000", data.Session, "check session")
	assert.Equal(t, "vvvv", data.Answer, "check answer")
}
