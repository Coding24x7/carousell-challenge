package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/", "public/swagger/")
	Files("/index.html", "public/swagger/")
	Files("/swagger", "public/swagger/")

})
