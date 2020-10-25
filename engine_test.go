package mx_test

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/aslrousta/mx"
)

func ExampleEngine_copywriter() {
	b := bytes.Buffer{}
	e := mx.Engine{
		Reader: strings.NewReader("Hello, World"),
		Writer: &b,
	}
	if err := e.Execute(); err != nil {
		panic(err)
	}
	fmt.Println(b.String())
	// Output: Hello, World
}

func ExampleEngine_unicode() {
	b := bytes.Buffer{}
	e := mx.Engine{
		Reader: strings.NewReader("こんにちは世界"),
		Writer: &b,
	}
	if err := e.Execute(); err != nil {
		panic(err)
	}
	fmt.Println(b.String())
	// Output: こんにちは世界
}
