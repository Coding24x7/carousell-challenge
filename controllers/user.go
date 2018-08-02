package controllers

import (
	"github.com/Coding24x7/carousell-challenge/app"
	"github.com/goadesign/goa"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here

	return nil
	// UserController_Create: end_implement
}

// Get runs the get action.
func (c *UserController) Get(ctx *app.GetUserContext) error {
	// UserController_Get: start_implement

	// Put your logic here

	return nil
	// UserController_Get: end_implement
}

// Remove runs the remove action.
func (c *UserController) Remove(ctx *app.RemoveUserContext) error {
	// UserController_Remove: start_implement

	// Put your logic here

	return nil
	// UserController_Remove: end_implement
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here

	return nil
	// UserController_Show: end_implement
}
