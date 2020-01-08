package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/owlsn/apis/src/datamodels"
)

type QueryPost func(table string) bool

type PostRepository interface {
	GetAll(m map[string]string) (*datamodels.Post, error)
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostMysqlRepository{source: db}
}

type PostMysqlRepository struct {
	source *gorm.DB
}

func (r *PostMysqlRepository) GetAll(m map[string]string) (*datamodels.Post, error) {
	for k, v := range m {
		r.source.Where(k, v)
	}

	ret := &datamodels.Post{}
	if err := r.source.First(ret, "id = ?", 1).Error; err != nil {
		return nil, err
	}

	return ret, nil
}
