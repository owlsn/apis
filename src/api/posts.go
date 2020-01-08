package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/owlsn/apis/src/services"
	"github.com/owlsn/apis/src/utils/json"
)

type PostsController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	Service services.PostService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

func (c *PostsController) GetAll() string {
	post, err := c.Service.GetAll()
	if err != nil{
		return "get post failed"
	}else{
		return post.Title
	}
}

func (c *PostsController) GetOne() *json.JsonResult {
	post, err := c.Service.GetAll()
	if err != nil{
		return json.JsonErrorData(-1, "get failed", err)
	}else{
		return json.JsonData(post)
	}
}

func (c *PostsController) Get() string {
	return "entry posts"
}