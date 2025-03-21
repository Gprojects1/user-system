package repository

import (
	"user-system/pkg/errors"
	"user-system/pkg/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *model.User) (*model.User, error)
	FindById(id uint) (*model.User, error)
	DeleteById(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Save(review *model.User) (*model.User, error) {
	if err := u.db.Save(review).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			err = errors.Wrapf(err, errors.NotSaved.Message())
			err = errors.AddErrorContext(err, "id", "can't save record with this id")
		default:
			err = errors.Newf(errors.NotSaved.Message())
		}
		return nil, err
	}
	return review, nil
}

func (u *userRepository) FindById(id uint) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			err = errors.Wrapf(err, errors.NotFound.Message())
			err = errors.AddErrorContext(err, "id", "can't found record with this id")
		default:
			err = errors.Newf(errors.NotFound.Message())
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) DeleteById(id uint) error {
	return u.db.Where("id = ?", id).Delete(new(model.User)).Error
}
