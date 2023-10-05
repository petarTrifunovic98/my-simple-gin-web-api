package services

import (
	"my-simple-gin-web-api/entities"
	"my-simple-gin-web-api/repository"
)

type UserService interface {
	Save(entities.User) entities.User
	FindAll() []entities.User
	FindByID(id any) entities.User
}

type userService struct {
	userRepository repository.UserRepository
}

func New(repository repository.UserRepository) UserService {
	return &userService{
		userRepository: repository,
	}
}

func (service *userService) Save(user entities.User) entities.User {
	service.userRepository.Save(user)
	return user
}

func (service *userService) FindAll() []entities.User {
	return service.userRepository.FindAll()
}

func (service *userService) FindByID(id any) entities.User {
	return service.userRepository.FindByID(id)
}
