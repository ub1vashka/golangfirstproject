package storageerror

import "errors"

var (
	ErrBookAlredyExist  = errors.New("book alredy exist")
	ErrBookByIdNotFound = errors.New("book id not found")
	ErrEmptyStorage     = errors.New("book storage is empty")
	ErrBookNoFound      = errors.New("book not found")
	ErrBookIDNotFound   = errors.New("book id not found")
)
