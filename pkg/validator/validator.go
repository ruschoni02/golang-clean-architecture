package validator

import (
	"errors"
	"reflect"
	"sync"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_loc "github.com/go-playground/locales/en"
	en_trans "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	once     sync.Once
	validate *validator.Validate
	trans    ut.Translator
	tagname  string
}

func New(tagname string) *Validator {
	return &Validator{
		tagname: tagname,
	}
}

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *Validator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			e, ok := err.(validator.ValidationErrors)
			if !ok {
				return err
			}

			return errors.New(e[0].Translate(v.trans))
		}
	}
	return nil
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://godoc.org/gopkg.in/go-playground/validator.v8
func (v *Validator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *Validator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()

		en := en_loc.New()
		uni := ut.New(en, en)
		v.trans, _ = uni.GetTranslator("en")
		en_trans.RegisterDefaultTranslations(v.validate, v.trans)

		v.validate.SetTagName(v.tagname)
	})
}
