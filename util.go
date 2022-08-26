package blut

import (
	"github.com/pkg/errors"
)

func okOrError(err error, s string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return errors.Wrapf(err, s, args...)
}

func i2b(i interface{}) []byte {
	switch v := i.(type) {

	}
}
