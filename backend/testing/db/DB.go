//testing
package db

import (
	"github.com/pocket-id/pocket-id/backend/internal/model"
	"errors"
//	"fmt"
	//	"gorm.io/gorm"
//	"log"
)

type WrappedError struct {
	Context string
	Err     error
}

func (w *WrappedError) Error() string {
	return ""
}

func Wrap(err error, info string) *WrappedError {
	return &WrappedError{
		Context: info,
		Err:     err,
	}
}

type DB struct {
    db string
}

func NewDBService() (*DB, error) {
	return &DB {
		db:               "dummy",
	}, nil
}


func (db *DB) First(user *model.User, query string, recipientUserId string) error {
	user.Email="alexlehm@tilde.green"
	user.FirstName="Alexander"
	user.LastName="L."

	err := errors.New("boom!")
	err = Wrap(err, "main")
	
	return err
}

