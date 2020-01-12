package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/owlsn/apis/src/services"
	"github.com/owlsn/apis/src/utils/json"
)

type UserController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	Service services.UserService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

func (c *UserController) GetAll() string {
	user, err := c.Service.GetAll()
	if err != nil {
		return "get user failed"
	} else {
		return user.Username
	}
}

func (c *UserController) GetOne() *json.JsonResult {
	user, err := c.Service.GetAll()
	if err != nil {
		return json.JsonErrorData(-1, "get failed", err)
	} else {
		return json.JsonData(user)
	}
}

func (c *UserController) GetLogin() *json.JsonResult {
	var (
		username = c.Ctx.PostValueTrim("username")
		password = c.Ctx.PostValueTrim("password")
	)
	exist, err := c.Service.GetAll()
	if err != nil {
		return json.JsonErrorData(-1, "get failed", err)
	} else {
		return json.JsonData(user)
	}
}

func (c *UserController) Get() string {
	return "entry users"
}
