package validation

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/cjtoolkit/validate/vString"
	"github.com/gorilla/schema"
	"github.com/thedevsaddam/govalidator"
)

func BenchmarkValidationSimple(b *testing.B) {
	rules := []vString.ValidationRule{vString.Mandatory(), vString.BetweenRune(2, 4)}
	params := url.Values{}
	params.Add("username", "ttt")

	for n := 0; n < b.N; n++ {
		vString.Validate(params.Get("username"), rules...)
	}
}

func BenchmarkGoValidatorSimple(b *testing.B) {
	params := url.Values{}
	params.Add("username", "ttt")
	r, _ := http.NewRequest("GET", "/hello", nil)
	r.Form = params
	rulesList := govalidator.MapData{
		"username": []string{"required", "between:2,4"},
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

func BenchmarkSchemaSimple(b *testing.B) {
	type User struct {
		Username string `schema:"username,required"`
	}

	params := url.Values{}
	params.Add("username", "ttt")

	decoder := schema.NewDecoder()

	for n := 0; n < b.N; n++ {
		var user User
		decoder.Decode(&user, params)
	}
}

func BenchmarkValidationSimpleWithError(b *testing.B) {
	rules := []vString.ValidationRule{vString.Mandatory(), vString.BetweenRune(2, 4)}
	params := url.Values{}
	params.Add("username", "")

	for n := 0; n < b.N; n++ {
		vString.Validate(params.Get("username"), rules...)
	}
}

func BenchmarkGoValidatorSimpleWithError(b *testing.B) {
	params := url.Values{}
	params.Add("username", "")
	r, _ := http.NewRequest("GET", "/hello", nil)
	r.Form = params
	rulesList := govalidator.MapData{
		"username": []string{"required", "between:2,4"},
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

func BenchmarkSchemaSimpleWithError(b *testing.B) {
	type User struct {
		Username string `schema:"username,required"`
	}

	params := url.Values{}
	params.Add("username", "")

	decoder := schema.NewDecoder()

	for n := 0; n < b.N; n++ {
		var user User
		decoder.Decode(&user, params)
	}
}
