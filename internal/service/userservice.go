package service

import (
	"time"

	"github.com/ub1vashka/golangfirstproject/internal/domain/models"
	"github.com/ub1vashka/golangfirstproject/internal/logger"
)

type UserStorage interface {
	SaveUser(models.User) (string, error)
	ValidateUser(models.UserLogin) (string, error)
	GetUsers() ([]models.User, error)
	GetUser(string) (models.User, error)
	DeleteUser(string) error
}

type UserService struct {
	stor UserStorage
}

func NewUserService(stor UserStorage) UserService {
	return UserService{stor: stor}
}

func (us *UserService) GetUsers() ([]models.User, error) {
	return us.stor.GetUsers()
}

func (us *UserService) GetUser(uid string) (models.User, error) {
	return us.stor.GetUser(uid)
}

func (us *UserService) DeleteUser(uid string) error {
	return us.stor.DeleteUser(uid)
}

func (us *UserService) LoginUser(user models.UserLogin) (string, error) {
	log := logger.Get()
	uid, err := us.stor.ValidateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("validate user failed")
		return ``, err
	}
	return uid, nil
}

func (us *UserService) RegisterUser(user models.User) (string, error) {
	log := logger.Get()
	user.RegisterDate = time.Now()
	uid, err := us.stor.SaveUser(user)
	if err != nil {
		log.Error().Err(err).Msg("save user failed")
		return ``, err
	}
	return uid, nil
}
