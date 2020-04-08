package pkg

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	// Trans 单例翻译器
	Trans    ut.Translator
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		z := zh.New()
		uni := ut.New(z, z)

		Trans, _ = uni.GetTranslator("zh")
		zh_translations.RegisterDefaultTranslations(v, Trans)
	}
}
