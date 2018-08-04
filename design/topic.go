package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("topic", func() {
	BasePath("/topics")

	Action("show", func() {
		Routing(GET(""))
		Description("Show top topics")

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})

	Action("post", func() {
		Routing(POST(""))
		Description("Post new topic")

		Payload(func() {
			Attribute("content", String, "Content of the topic", func() {
				MaxLength(255)
			})
			Attribute("userName", String, "Author of this topic")
			Required("content", "userName")

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
			Attribute("userName", String, "username")
			Attribute("vote", String, "upvote/downvote topic", func() {
				Enum("up", "down")
			})

			Required("userName", "vote")
		})

		Response(OK, Any)
		Response(NotFound, String)
		Response(BadRequest, String)
		Response(Unauthorized, String)
		Response(InternalServerError, String)
	})
})
