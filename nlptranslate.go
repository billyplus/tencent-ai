package ai

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
)

const (
	URLTextTranslate = "https://api.ai.qq.com/fcgi-bin/nlp/nlp_texttranslate"
)

type TextLanguage int

const (
	// LangChinese 中文
	LangChinese TextLanguage = iota
	// LangEnglish 英文
	LangEnglish
	// LangJapanese 日文
	LangJapanese
	// LangKorean 韩文
	LangKorean
	// LangFrench 法文
	LangFrench
	// LangSpanish 西班牙文
	LangSpanish
	// LangItalian 意大利文
	LangItalian
	// LangGerman 德文
	LangGerman
	// LangTurkish 土耳其文
	LangTurkish
	// LangRussian 俄文
	LangRussian
	// LangPortuguese 葡萄牙文
	LangPortuguese
	// LangVietnamese 越南文
	LangVietnamese
	// LangIndonesian 印度尼西亚文
	LangIndonesian
	// LangMalaysian 马来西亚文
	LangMalaysian
	// LangThai 泰文
	LangThai
	// LangAuto 自动识别（中英互译）
	LangAuto
)

var (
	languageStr = []string{
		// LangChinese 中文
		LangChinese: "zh",
		// LangEnglish 英文
		LangEnglish: "en",
		// LangJapanese 日文
		LangJapanese: "jp",
		// LangKorean 韩文
		LangKorean: "kr",
		// LangFrench 法文
		LangFrench: "fr",
		// LangSpanish 西班牙文
		LangSpanish: "es",
		// LangItalian 意大利文
		LangItalian: "it",
		// LangGerman 德文
		LangGerman: "de",
		// LangTurkish 土耳其文
		LangTurkish: "tr",
		// LangRussian 俄文
		LangRussian: "ru",
		// LangPortuguese 葡萄牙文
		LangPortuguese: "pt",
		// LangVietnamese 越南文
		LangVietnamese: "vi",
		// LangIndonesian 印度尼西亚文
		LangIndonesian: "id",
		// LangMalaysian 马来西亚文
		LangMalaysian: "ms",
		// LangThai 泰文
		LangThai: "th",
		// LangAuto 自动识别（中英互译）
		LangAuto: "auto",
	}
)

var (
	langList = [][]TextLanguage{
		LangChinese:    []TextLanguage{LangEnglish, LangFrench, LangSpanish, LangItalian, LangGerman, LangTurkish, LangRussian, LangPortuguese, LangVietnamese, LangIndonesian, LangMalaysian, LangThai, LangJapanese, LangKorean},
		LangEnglish:    []TextLanguage{LangChinese, LangFrench, LangSpanish, LangItalian, LangGerman, LangTurkish, LangRussian, LangPortuguese, LangVietnamese, LangIndonesian, LangMalaysian, LangThai},
		LangFrench:     []TextLanguage{LangEnglish, LangChinese, LangSpanish, LangItalian, LangGerman, LangTurkish, LangRussian, LangPortuguese},
		LangSpanish:    []TextLanguage{LangEnglish, LangChinese, LangFrench, LangItalian, LangGerman, LangTurkish, LangRussian, LangPortuguese},
		LangItalian:    []TextLanguage{LangEnglish, LangChinese, LangFrench, LangSpanish, LangGerman, LangTurkish, LangRussian, LangPortuguese},
		LangGerman:     []TextLanguage{LangEnglish, LangChinese, LangFrench, LangSpanish, LangItalian, LangTurkish, LangRussian, LangPortuguese},
		LangTurkish:    []TextLanguage{LangEnglish, LangChinese, LangFrench, LangSpanish, LangItalian, LangGerman, LangRussian, LangPortuguese},
		LangRussian:    []TextLanguage{LangEnglish, LangChinese, LangFrench, LangSpanish, LangItalian, LangGerman, LangTurkish, LangPortuguese},
		LangPortuguese: []TextLanguage{LangEnglish, LangChinese, LangFrench, LangSpanish, LangItalian, LangGerman, LangTurkish, LangRussian},
		LangVietnamese: []TextLanguage{LangEnglish, LangChinese},
		LangIndonesian: []TextLanguage{LangEnglish, LangChinese},
		LangMalaysian:  []TextLanguage{LangEnglish, LangChinese},
		LangThai:       []TextLanguage{LangEnglish, LangChinese},
		LangJapanese:   []TextLanguage{LangChinese},
		LangKorean:     []TextLanguage{LangChinese},
	}
	langSupport []uint64
)

func init() {
	langSupport = make([]uint64, len(langList))
	for lang, lst := range langList {
		v := uint64(0)
		for _, l := range lst {
			v = v | 1<<l
		}
		langSupport[lang] = v
	}
}

type NLPTextTranslateData struct {
	SourceText string `json:"source_text"`
	TargetText string `json:"target_text"`
}

type NLPTextTranslateResult struct {
	Return  int                   `json:"ret"`
	Message string                `json:"msg"`
	Data    *NLPTextTranslateData `json:"data"`
}

// NLPTextTranslate 文本翻译
func (a *App) NLPTextTranslate(source TextLanguage, target TextLanguage, text string) (*NLPTextTranslateData, error) {
	v := langSupport[source]
	if v|uint64(target) == 0 {
		return nil, errors.Errorf("unsupport translation from %s to %s", languageStr[source], languageStr[target])
	}
	param := url.Values{}
	param.Add(KeyText, text)
	param.Add(KeySource, languageStr[source])
	param.Add(KeyTarget, languageStr[target])
	a.prepareRequestParam(&param)
	var result NLPTextTranslateResult
	fmt.Println(param.Encode())

	if err := a.do(URLTextTranslate, &param, &result); err != nil {
		return nil, errors.Wrap(err, "do request")
	}

	if result.Return != 0 {
		return nil, errors.Errorf("server return error=%d msg=%s", result.Return, result.Message)
	}
	return result.Data, nil
}
