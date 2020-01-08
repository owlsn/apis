package services

import (
	"github.com/owlsn/apis/src/repositories"
	"github.com/owlsn/apis/src/datamodels"
)

type PostService interface {
	GetAll() (*datamodels.Post, error) 
}

func NewPostService (repo repositories.PostRepository) PostService {
	return &postService{
		repo : repo,
	}
}

type postService struct {
	repo repositories.PostRepository
}

func (s *postService) GetAll() (*datamodels.Post, error ){
	var where map[string] string 
	where = make(map[string] string )
	where["1 = ?"] = "1"
	post, err := s.repo.GetAll(where)
	if err != nil{
		return nil, err
	}
	return post, nil
}