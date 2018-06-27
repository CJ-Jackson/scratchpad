package ctx

import (
	"testing"

	"fmt"

	"github.com/cjtoolkit/ctx"
)

type Complex interface {
	GetStr() string
	GetNumber() int64
}

type complex struct {
	Str    string
	Number int64
}

func (c complex) GetStr() string {
	return c.Str
}

func (c complex) GetNumber() int64 {
	return c.Number
}

func getComplexType(context ctx.BackgroundContext) Complex {
	type complexContext struct{}
	return context.Persist(complexContext{}, func() (interface{}, error) {
		return Complex(complex{
			Str:    "Hello World",
			Number: 5,
		}), nil
	}).(Complex)
}

func BenchmarkCtx(b *testing.B) {
	fmt.Printf("%#v", complex{
		Str:    "Hello World",
		Number: 5,
	})

	context := ctx.NewBackgroundContext()
	getComplexType(context)

	for n := 0; n < b.N; n++ {
		getComplexType(context)
	}
}
