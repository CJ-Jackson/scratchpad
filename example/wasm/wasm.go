package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"syscall/js"
)

func JQuery(args ...interface{}) js.Value {
	return js.Global().Call("jQuery", args...)
}

const htmlTemplate = `
<h2>{{ .Forename }} {{ .Surname }}</h2>
`

func checkAndUpdate(event js.Value, tpl *template.Template) {
	type context struct {
		Forename string
		Surname  string
	}
	buf := &bytes.Buffer{}
	tpl.Execute(buf, context{
		Forename: JQuery("#forename").Call("val").String(),
		Surname:  JQuery("#surname").Call("val").String(),
	})

	JQuery("#box").Call("html", buf.String())
}

func ipAddress() {
	res, _ := http.Get("https://icanhazip.com/")
	b, _ := ioutil.ReadAll(res.Body)

	JQuery("#ipAddress").Call("text", strings.TrimSpace(string(b)))
}

func main() {
	tpl := template.Must(template.New("template").Parse(htmlTemplate))

	JQuery("#forename, #surname").
		Call("on", "keyup", js.NewEventCallback(js.PreventDefault, func(event js.Value) {
			checkAndUpdate(event, tpl)
		}))

	ipAddress()

	select {}
}
