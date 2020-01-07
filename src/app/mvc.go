package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"

	"github.com/kataras/iris/v12/mvc"
	"github.com/owlsn/apis/src/api"
	"github.com/sirupsen/logrus"

	"github.com/owlsn/apis/src/datasources"
	"github.com/owlsn/apis/src/repositories"
	"github.com/owlsn/apis/src/services"
	// "github.com/owlsn/apis/src/middleware"

	"time"
)

// MVC : mvc func
func MVC(app *mvc.Application) {
	
	// You can use normal middlewares at MVC apps of course.
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	// Register dependencies which will be binding to the controller(s),
	// can be either a function which accepts an iris.Context and returns a single value (dynamic binding)
	// or a static struct value (service).
	app.Register(
		sessions.New(sessions.Config{}).Start,
		&api.PrefixedLogger{Prefix: "DEV"},
	)

	// GET: http://localhost:8080/basic
	// GET: http://localhost:8080/basic/custom
	// GET: http://localhost:8080/basic/custom2
	app.Handle(new(api.BasicController))

	// All dependencies of the parent *mvc.Application
	// are cloned to this new child,
	// thefore it has access to the same session as well.
	// GET: http://localhost:8080/basic/sub
	app.Party("/sub").
		Handle(new(api.BasicSubController))
}

func User(user *mvc.Application) {
	// ---- Serve our controllers. ----
	// "/user" based mvc application.
	// Prepare our repositories and services.
	db, err := datasource.LoadUsers(datasource.Memory)
	if err != nil {
		logrus.Error("error while loading the users: %v")
		return
	}
	repo := repositories.NewUserRepository(db)
	userService := services.NewUserService(repo)

	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookiename",
		Expires: 24 * time.Hour,
	})
	
	user.Register(
		userService,
		sessManager.Start,
	)
	user.Handle(new(api.UserController))

}

func Users(users *mvc.Application){
	// Prepare our repositories and services.
	logrus.Info("entry users")
	db, err := datasource.LoadUsers(datasource.Memory)
	if err != nil {
		logrus.Error("error while loading the users: %v")
		return
	}
	logrus.Info(db)
	repo := repositories.NewUserRepository(db)
	userService := services.NewUserService(repo)

	// Add the basic authentication(admin:password) middleware
	// for the /users based requests.
	// users.Router.Use(middleware.BasicAuth)
	// Bind the "userService" to the UserController's Service (interface) field.
	users.Register(userService)
	users.Handle(new(api.UsersController))

}