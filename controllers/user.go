package controllers

import (
	"github.com/Coding24x7/carousell-challenge/app"
	"github.com/Coding24x7/carousell-challenge/lib"
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

// Register runs the create action.
func (c *UserController) Register(ctx *app.RegisterUserContext) error {
	u, err := lib.RegisterUser(ctx.Payload.Name, ctx.Payload.Password, ctx.Payload.Country)

	if err != nil {
		return processErr(ctx, err)
	}
	return ctx.OK(u)
}

// Login runs the get action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	u, err := lib.LoginUser(ctx.UserName, ctx.Password)
	if err != nil {
		return processErr(ctx, err)
	}

	return ctx.OK(u)
}

// Remove runs the remove action.
func (c *UserController) Remove(ctx *app.RemoveUserContext) error {
	u, err := lib.DeleteUser(ctx.UserName)
	if err != nil {
		return processErr(ctx, err)
	}

	return ctx.OK(u)
}
