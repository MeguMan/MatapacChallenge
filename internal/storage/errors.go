package storage

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
)

var ErrUniqueKeyViolation = errors.New("unique key violation")

func handleErr(err error) error {
	pgErr, ok := err.(*pq.Error)
	fmt.Println("CODE: ", pgErr.Code)
	if ok && pgErr.Code == "23505" {
		return fmt.Errorf("%w: %v", ErrUniqueKeyViolation, err)
	}

	return err
}
