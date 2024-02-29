package utils

import (
	"fmt"
)

func AppendErr(err, newErr error) error {
	if err == nil {
		return newErr
	}
	return fmt.Errorf("%v, %w", err, newErr)
}
