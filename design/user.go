package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("user", func() {
	BasePath("/users")

	Action("show", func() {
		Routing(GET(""))
		Description("Get all users")

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})

	Action("create", func() {
		Routing(POST(""))
		Description("Create new user")

		Payload(func() {
			Attribute("name", String, "name of the user")
			Attribute("country", String, "country of the user")
			Required("name", "country")

		})

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})

	Action("get", func() {
		Routing(GET("/:userName"))
		Description("Get user details by name")

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})

	Action("remove", func() {
		Description("remove user")
		Routing(DELETE("/:userID"))

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})
})
