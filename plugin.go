package main // import "github.com/juhwany/go_plugin"

import "github.com/juhwany/go_server/v2/apis"

func Foo(bar apis.Bar) {
	bar.DoSomething()
}
