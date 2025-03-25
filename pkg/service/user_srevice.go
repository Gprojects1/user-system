package service

import (
	"strconv"
	"user-system/pkg/dto"
	"user-system/pkg/errors"
	"user-system/pkg/model"
	"user-system/pkg/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (s *UserService) DeleteUserById(idString string) error {

	id, err := strconv.Atoi(idString)
	if err != nil {
		err = errors.BadRequest.Wrapf(err, errors.WrongType.Message())
		err = errors.AddErrorContext(err, "id", "wrong id format, should be an integer")

		return err
	}
	err = s.UserRepo.DeleteById(uint(id))
	return err
}

func (s *UserService) UpdateUser(idString string, req *dto.UpdateUserRequest) (*model.User, error) {

	id, err := strconv.Atoi(idString)
	if err != nil {
		err = errors.BadRequest.Wrapf(err, errors.WrongType.Message())
		err = errors.AddErrorContext(err, "id", "wrong id format, should be an integer")

		return &model.User{}, err
	}

	exUser, err := s.UserRepo.FindById(uint(id))
	if err != nil {
		return &model.User{}, err
	}
	user := &model.User{Email: req.Email, Role: exUser.Role, Password: exUser.Password, Model: exUser.Model}
	if err = s.UserRepo.DeleteById(uint(id)); err != nil {
		return &model.User{}, err
	}
	return s.UserRepo.Save(user)
}

func (s *UserService) AddUser(user *model.User) (*model.User, error) {

	if user.Email == "" || user.Role == "" {
		err := errors.EmptyData.Newf(errors.EmptyData.Message())
		err = errors.AddErrorContext(err, "id", "Email and Role can't be empty")
		return nil, err
	}
	return s.UserRepo.Save(user)
}
