package services

import (
	"github.com/owlsn/apis/src/datamodels"
	"github.com/owlsn/apis/src/repositories"
	"github.com/owlsn/apis/src/utils/database"
)

type UserService interface {
	GetAll() (*datamodels.User, error)
	Exist(username string, password string) (bool, error)
}

func NewUserService(repo repositories.AuthRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositories.AuthRepository
}

func (s *userService) GetAll() (*datamodels.User, error) {
	cnd := database.NewSqlCnd()
	var where map[string]string
	where = make(map[string]string)
	where["1 = ?"] = "1"
	user, err := s.repo.GetAll(where)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s userService) Exist(username string, password string) (bool, error) {
	var where map[string]string
	where = make(map[string]string)
	where["username"]
}
