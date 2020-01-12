package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/owlsn/apis/src/datamodels"
)

type AuthRepository interface {
	GetAll(m map[string] string) (*datamodels.User, error)
	GetOne(m map[string] string) (bool, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthUserRepository{source: db}
}

type AuthUserRepository struct {
	source *gorm.DB
}

func (r *AuthUserRepository) GetAll(m map[string]string) (*datamodels.User, error) {
	for k, v := range m {
		r.source.Where(k, v)
	}

	ret := &datamodels.User{}
	if err := r.source.First(ret, "id = ?", 1).Error; err != nil {
		return nil, err
	}

	return ret, nil
}

func (r *AuthUserRepository) GetOne(m map[string]string) (bool, error) {
	
}
