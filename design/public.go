package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/*filepath", "public/")
})
