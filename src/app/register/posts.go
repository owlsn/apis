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
func Posts(posts *mvc.Application){
	logrus.Info("register posts")
	db, err := datasource.Load(datasource.MySQL)
	if err != nil {
		logrus.Errorf("error while loading the users: ", err.Error())
		return
	}
	repo := repositories.NewPostRepository(db)
	postService := services.NewPostService(repo)
	posts.Register(postService)
	posts.Handle(new(api.PostsController))
}