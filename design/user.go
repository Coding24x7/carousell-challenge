package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("user", func() {
	BasePath("/users")

	Action("register", func() {
		Routing(POST(""))
		Description("Register new user")

		Payload(func() {
			Attribute("name", String, "username")
			Attribute("password", String, "password of the user")
			Attribute("country", String, "country of the user")
			Required("name", "password", "country")

		})

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})

	Action("login", func() {
		Routing(GET("/:userName"))
		Description("Login user details by name")

		Params(func() {
			Attribute("password", String, "password of the user")
			Required("password")
		})

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})

	Action("remove", func() {
		Description("remove user")
		Routing(DELETE("/:userName"))

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})
})
