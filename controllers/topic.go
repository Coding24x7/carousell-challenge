package controllers

import (
	"fmt"

	"github.com/Coding24x7/twitter-clone/app"
	"github.com/Coding24x7/twitter-clone/lib"
	"github.com/goadesign/goa"
)

// TopicController implements the topic resource.
type TopicController struct {
	*goa.Controller
}

// NewTopicController creates a topic controller.
func NewTopicController(service *goa.Service) *TopicController {
	return &TopicController{Controller: service.NewController("TopicController")}
}

// Post runs the post action.
func (c *TopicController) Post(ctx *app.PostTopicContext) error {
	t, err := lib.PostTopic(ctx.Payload.Author, ctx.Payload.Content)

	if err != nil {
		return processErr(ctx, err)
	}
	return ctx.OK(t)
}

// Show runs the show action.
func (c *TopicController) Show(ctx *app.ShowTopicContext) error {
	topics := lib.ShowTopics()
	return ctx.OK(topics)
}

// Vote runs the vote action.
func (c *TopicController) Vote(ctx *app.VoteTopicContext) error {
	if ctx.Payload.Vote == "up" {
		err := lib.UpvoteTopic(ctx.TopicID, ctx.Payload.UserName)
		if err != nil {
			return processErr(ctx, err)
		}
		return ctx.OK(nil)

	} else if ctx.Payload.Vote == "down" {
		err := lib.DownvoteTopic(ctx.TopicID, ctx.Payload.UserName)
		if err != nil {
			return processErr(ctx, err)
		}
		return ctx.OK(nil)

	}
	return ctx.BadRequest(fmt.Sprintf("wrong parameter vote %s", ctx.Payload.Vote))
}
