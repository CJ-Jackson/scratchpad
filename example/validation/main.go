package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/cjtoolkit/validate/vError"
	"github.com/cjtoolkit/validate/vInt"
	"github.com/cjtoolkit/validate/vString"
	"github.com/julienschmidt/httprouter"
)

// Easy to mentalise with HTML template
type Form struct {
	Checked         bool
	Valid           bool
	Name            string
	NameErr         error
	Email           string
	EmailErr        error
	EmailConfirm    string
	EmailConfirmErr error
	Age             int64
	AgeErr          error
}

// Template for the Form struct above.
const FormHtml = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css" integrity="sha384-9gVQ4dYFwwWSjIDZnLEWnxCjeSWFphJiwGPXr1jddIhOegiu1FwO5qRGvFXOdJZ4" crossorigin="anonymous">

    <title>Form Test</title>
  </head>
  <body>
	<div class="container">
	<h1>Form Test</h1>

	{{ if .Checked }}
		{{ if .Valid }}
			<div class="alert alert-success">Success</div>
		{{ else }}
			<div class="alert alert-danger">Fail</div>
		{{ end }}
	{{ end }}

	<form method="post" novalidate>

		<div class="form-group">
			<label for="name">Name:</label>
			<input id="name" type="text"
				class="form-control {{ if .Checked }}{{ if not .NameErr }}is-valid{{ else }}is-invalid{{ end }}{{ end }}" name="name" value="{{ .Name }}">
			{{ if .Checked }}{{ if .NameErr }}<div class="invalid-feedback">{{ .NameErr }}</div>{{ end }}{{ end }}
		</div>

		<div class="form-group">
			<label for="email">Email:</label>
			<input id="email" type="email"
				class="form-control {{ if .Checked }}{{ if not .EmailErr }}is-valid{{ else }}is-invalid{{ end }}{{ end }}" name="email" value="{{ .Email }}">
			{{ if .Checked }}{{ if .EmailErr }}<div class="invalid-feedback">{{ .EmailErr }}</div>{{ end }}{{ end }}
		</div>

		<div class="form-group">
			<label for="email-confirm">Confirm Email:</label>
			<input id="email-confirm" type="email"
				class="form-control {{ if .Checked }}{{ if not .EmailConfirmErr }}is-valid{{ else }}is-invalid{{ end }}{{ end }}" name="email-confirm" value="{{ .EmailConfirm }}">
			{{ if .Checked }}{{ if .EmailConfirmErr }}<div class="invalid-feedback">{{ .EmailConfirmErr }}</div>{{ end }}{{ end }}
		</div>

		<div class="form-group">
			<label for="age">Age:</label>
			<input id="age" type="email"
				class="form-control {{ if .Checked }}{{ if not .AgeErr }}is-valid{{ else }}is-invalid{{ end }}{{ end }}" name="age" value="{{ .Age }}">
			{{ if .Checked }}{{ if .AgeErr }}<div class="invalid-feedback">{{ .AgeErr }}</div>{{ end }}{{ end }}
		</div>

		<input class="form-control" type="submit">
	</form>

	</div>
   
    <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.0/umd/popper.min.js" integrity="sha384-cs/chFZiN24E4KMATLdqdvsezGxaGsi4hLGOzlXwp5UZB1LY//20VyM2taTB4QvJ" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.0/js/bootstrap.min.js" integrity="sha384-uefMccjFJAIv6A+rW+L4AHf99KvxDjWSu1z9VI8SKNVmz4sk7buKt/6v9KI65qnm" crossorigin="anonymous"></script>
  </body>
</html>`

func newForm() *Form {
	return &Form{Valid: true}
}

// Form Builder Validator
type FormBuilderValidator struct {
	nameRules  []vString.ValidationRule
	emailRules []vString.ValidationRule
	ageRules   []vInt.ValidationRule
}

func NewFormBuilderValidator() FormBuilderValidator {
	return FormBuilderValidator{
		nameRules:  []vString.ValidationRule{vString.Mandatory(), vString.BetweenRune(5, 50)},
		emailRules: []vString.ValidationRule{vString.Mandatory(), vString.Email()},
		ageRules:   []vInt.ValidationRule{vInt.Mandatory(), vInt.Min(21)},
	}
}

func (v FormBuilderValidator) NewBlankForm() Form {
	f := newForm()
	return *f
}

func (v FormBuilderValidator) NewValidatedForm(params url.Values) Form {
	f := newForm()
	f.Checked = true

	f.Name, f.NameErr = vString.Validate(params.Get("name"), v.nameRules...)
	f.Email, f.EmailErr = vString.Validate(params.Get("email"), v.emailRules...)
	f.EmailConfirm, f.EmailConfirmErr = vString.Validate(params.Get("email-confirm"), append(v.emailRules,
		vString.MustMatch(f.Email, "Email"))...)
	f.Age, f.AgeErr = vInt.ValidateFromString(params.Get("age"), v.ageRules...)

	f.Valid = vError.CheckErr(f.NameErr, f.EmailErr, f.EmailConfirmErr, f.AgeErr)
	return *f
}

func main() {
	router := httprouter.New()
	formBuilderValidator := NewFormBuilderValidator()
	t := template.Must(template.New("form").Parse(FormHtml))

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		f := formBuilderValidator.NewBlankForm()
		t.Execute(writer, f)
	})

	router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		request.ParseForm()
		f := formBuilderValidator.NewValidatedForm(request.PostForm)
		t.Execute(writer, f)
	})

	fmt.Println("Running Server at :8080")
	log.Print(http.ListenAndServe(":8080", router))
}
