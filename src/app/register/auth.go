package register

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/owlsn/apis/src/api"
	"github.com/sirupsen/logrus"

	"github.com/owlsn/apis/src/datasources"
	"github.com/owlsn/apis/src/repositories"
	"github.com/owlsn/apis/src/services"
	// "github.com/owlsn/apis/src/middleware"
)

// Posts : Posts
func Auth(auth *mvc.Application) {
	logrus.Info("register auth")
	db, err := datasource.Load(datasource.MySQL)
	if err != nil {
		logrus.Errorf("error while loading the datasource: ", err.Error())
		return
	}
	repo := repositories.NewAuthRepository(db)
	userService := services.NewUserService(repo)
	auth.Register(userService)
	auth.Handle(new(api.UserController))
}
