package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("topic", func() {
	BasePath("/topics")

	Action("show", func() {
		Routing(GET(""))
		Description("Get top topics")

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})

	Action("create", func() {
		Routing(POST(""))
		Description("Create new topic")

		Payload(func() {
			Attribute("content", String, "Content of the topic")
			Attribute("userID", String, "Creator of this topic")
			Required("content", "userID")

		})

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})

	Action("vote", func() {
		Description("upvote/downvote cluster")
		Routing(PATCH("/:topicID"))

		Payload(func() {
			Attribute("userID", String, "user id")
			Attribute("vote", String, "upvote/downvote topic", func() {
				Enum("up", "down") // And validation rules
			})

			Required("userID", "vote")
		})

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})

	Action("delete", func() {
		Description("remove the topic")
		Routing(DELETE("/:topicID"))

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})
})
