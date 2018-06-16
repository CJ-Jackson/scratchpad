package validation

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/cjtoolkit/validate/vFloat"
	"github.com/cjtoolkit/validate/vInt"
	"github.com/cjtoolkit/validate/vString"
	"github.com/gorilla/schema"
	"github.com/thedevsaddam/govalidator"
)

func BenchmarkValidationComplex(b *testing.B) {
	userNameRules := []vString.ValidationRule{vString.Mandatory(), vString.BetweenRune(2, 4)}
	intRules := []vInt.ValidationRule{vInt.Between(4, 6)}
	floatRules := []vFloat.ValidationRule{vFloat.Between(5.4, 5.6)}
	emailRules := []vString.ValidationRule{vString.Email()}

	params := url.Values{}
	params.Add("username", "ttt")
	params.Add("int", "5")
	params.Add("float", "5.5")
	params.Add("email", "test@example.com")

	for n := 0; n < b.N; n++ {
		vString.Validate(params.Get("username"), userNameRules...)
		vInt.ValidateFromString(params.Get("int"), intRules...)
		vFloat.ValidateFromString(params.Get("float"), floatRules...)
		vString.Validate(params.Get("email"), emailRules...)
	}
}

func BenchmarkGoValidatorComplex(b *testing.B) {
	params := url.Values{}
	params.Add("username", "ttt")
	params.Add("int", "5")
	params.Add("float", "5.5")
	params.Add("email", "test@example.com")
	r, _ := http.NewRequest("GET", "/hello", nil)
	r.Form = params
	rulesList := govalidator.MapData{
		"username": []string{"required", "between:2,4"},
		"int":      []string{"between:4,6"},
		"float":    []string{"between:5.4,5.6"},
		"email":    []string{"required", "min:4", "max:20", "email"},
	}
	opts := govalidator.Options{
		Request: r,
		Rules:   rulesList,
	}
	v := govalidator.New(opts)
	for n := 0; n < b.N; n++ {
		v.Validate()
	}
}

func BenchmarkSchemaComplex(b *testing.B) {
	type User struct {
		Username string  `schema:"username,required"`
		Int      int64   `schema:"int,required"`
		Float    float64 `schema:"float,required"`
		Email    string  `schema:"email,required"`
	}

	params := url.Values{}
	params.Add("username", "ttt")
	params.Add("int", "5")
	params.Add("float", "5.5")
	params.Add("email", "test@example.com")

	decoder := schema.NewDecoder()

	for n := 0; n < b.N; n++ {
		var user User
		decoder.Decode(&user, params)
	}
}

func BenchmarkValidationComplexWithError(b *testing.B) {
	userNameRules := []vString.ValidationRule{vString.Mandatory(), vString.BetweenRune(2, 4)}
	intRules := []vInt.ValidationRule{vInt.Between(4, 6)}
	floatRules := []vFloat.ValidationRule{vFloat.Between(5.4, 5.6)}
	emailRules := []vString.ValidationRule{vString.Email()}

	params := url.Values{}
	params.Add("username", "")
	params.Add("int", "7")
	params.Add("float", "5.7")
	params.Add("email", "")

	for n := 0; n < b.N; n++ {
		vString.Validate(params.Get("username"), userNameRules...)
		vInt.ValidateFromString(params.Get("int"), intRules...)
		vFloat.ValidateFromString(params.Get("float"), floatRules...)
		vString.Validate(params.Get("email"), emailRules...)
	}
}

func BenchmarkGoValidatorComplexWithError(b *testing.B) {
	params := url.Values{}
	params.Add("username", "")
	params.Add("int", "7")
	params.Add("float", "5.7")
	params.Add("email", "")
	r, _ := http.NewRequest("GET", "/hello", nil)
	r.Form = params
	rulesList := govalidator.MapData{
		"username": []string{"required", "between:2,4"},
		"int":      []string{"between:4,6"},
		"float":    []string{"between:5.4,5.6"},
		"email":    []string{"required", "min:4", "max:20", "email"},
	}
	opts := govalidator.Options{
		Request: r,
		Rules:   rulesList,
	}
	v := govalidator.New(opts)
	for n := 0; n < b.N; n++ {
		v.Validate()
	}
}

func BenchmarkSchemaComplexWithError(b *testing.B) {
	type User struct {
		Username string  `schema:"username,required"`
		Int      int64   `schema:"int,required"`
		Float    float64 `schema:"float,required"`
		Email    string  `schema:"email,required"`
	}

	params := url.Values{}
	params.Add("username", "")
	params.Add("int", "")
	params.Add("float", "")
	params.Add("email", "")

	decoder := schema.NewDecoder()

	for n := 0; n < b.N; n++ {
		var user User
		decoder.Decode(&user, params)
	}
}
