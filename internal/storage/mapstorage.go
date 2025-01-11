package storage

import (
	"errors"

	"github.com/ub1vashka/golangfirstproject/internal/domain/models"
	"github.com/ub1vashka/golangfirstproject/internal/logger"
	"github.com/ub1vashka/golangfirstproject/internal/storage/storageerror"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type MapStorage struct {
	stor  map[string]models.User
	bStor map[string]models.Book
}

func New() *MapStorage {
	return &MapStorage{
		stor:  make(map[string]models.User),
		bStor: make(map[string]models.Book),
	}
}

func (ms *MapStorage) SaveUser(user models.User) (string, error) {
	log := logger.Get()
	for _, usr := range ms.stor {
		if user.Email == usr.Email {
			return ``, errors.New("user alredy exist")
		}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return ``, err
	}
	user.Password = string(hash)
	uid := uuid.New()
	user.UID = uid
	ms.stor[user.UID.String()] = user
	log.Debug().Any("storage", ms.stor).Msg("check storage")
	return uid.String(), nil
}

func (ms *MapStorage) DeleteUser(uid string) error {
	_, ok := ms.stor[uid]
	if !ok {
		return storageerror.ErrUserNotFound
	}
	delete(ms.stor, uid)
	return nil
}

func (ms *MapStorage) GetUsers() ([]models.User, error) {
	if len(ms.stor) == 0 {
		return nil, storageerror.ErrEmptyStorage
	}
	var users []models.User
	for _, user := range ms.stor {
		users = append(users, user)
	}
	return users, nil
}

func (ms *MapStorage) GetUser(uid string) (models.User, error) {
	user, ok := ms.stor[uid]
	if !ok {
		return models.User{}, storageerror.ErrUserNotFound
	}
	return user, nil
}

func (ms *MapStorage) ValidateUser(user models.UserLogin) (string, error) {
	for key, usr := range ms.stor {
		if user.Email == usr.Email {
			if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(user.Passoword)); err != nil {
				return ``, errors.New("invalid user password")
			}
			return key, nil
		}
	}
	return ``, errors.New("user no exist")
}

func (ms *MapStorage) SaveBook(book models.Book) (string, error) {
	log := logger.Get()
	for _, b := range ms.bStor {
		if book.Lable == b.Lable && book.Author == b.Author {
			return ``, storageerror.ErrBookAlredyExist
		}
	}
	bID := uuid.New()
	book.BID = bID
	ms.bStor[book.BID.String()] = book
	log.Debug().Any("book storage", ms.stor).Msg("check storage")
	return bID.String(), nil
}

func (ms *MapStorage) GetBooks() ([]models.Book, error) {
	if len(ms.bStor) == 0 {
		return nil, storageerror.ErrEmptyStorage
	}
	var books []models.Book
	for _, book := range ms.bStor {
		books = append(books, book)
	}
	return books, nil
}

func (ms *MapStorage) GetBook(bid string) (models.Book, error) {
	book, ok := ms.bStor[bid]
	if !ok {
		return models.Book{}, storageerror.ErrBookNoFound
	}
	return book, nil
}

func (ms *MapStorage) DeleteBook(bid string) error {
	_, ok := ms.bStor[bid]
	if !ok {
		return storageerror.ErrBookNoFound
	}
	delete(ms.bStor, bid)
	return nil
}
