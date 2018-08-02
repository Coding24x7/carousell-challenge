package controllers

import (
	"github.com/Coding24x7/carousell-challenge/app"
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

// Create runs the create action.
func (c *TopicController) Create(ctx *app.CreateTopicContext) error {
	// TopicController_Create: start_implement

	// Put your logic here

	return nil
	// TopicController_Create: end_implement
}

// Delete runs the delete action.
func (c *TopicController) Delete(ctx *app.DeleteTopicContext) error {
	// TopicController_Delete: start_implement

	// Put your logic here

	return nil
	// TopicController_Delete: end_implement
}

// Show runs the show action.
func (c *TopicController) Show(ctx *app.ShowTopicContext) error {
	// TopicController_Show: start_implement

	// Put your logic here

	return nil
	// TopicController_Show: end_implement
}

// Vote runs the vote action.
func (c *TopicController) Vote(ctx *app.VoteTopicContext) error {
	// TopicController_Vote: start_implement

	// Put your logic here

	return nil
	// TopicController_Vote: end_implement
}
