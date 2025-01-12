package storageerror

import "errors"

var (
	ErrUserNotFound     = errors.New("user id not found")
	ErrEmptyUserStorage = errors.New("user storage is empty")

	ErrBookAlredyExist = errors.New("book alredy exist")
	ErrEmptyStorage    = errors.New("book storage is empty")
	ErrBookNoFound     = errors.New("book not found")
	ErrBookIDNotFound  = errors.New("book id not found")
)
