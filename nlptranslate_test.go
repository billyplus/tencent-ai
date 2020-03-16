package ai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNLPTextTranslate(t *testing.T) {
	ResetDefaultApp(111111, "xxxxxxxxx")
	source := "你叫什么名字？？？？？？"
	data, err := app.NLPTextTranslate(LangChinese, LangEnglish, source)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, source, data.SourceText, "check source")
	assert.NotEmpty(t, data.TargetText, "check target")
}
