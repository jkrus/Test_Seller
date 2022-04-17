package errors

import "github.com/pkg/errors"

const (
	errLoadConfigMsg      = "load config"
	errOpenDatabaseMsg    = "open database"
	errStartHTTPServerMsg = "start http Server"
)

func ErrLoadConfig(w error) error {
	return errors.Wrap(w, errLoadConfigMsg)
}

func ErrOpenDatabase(w error) error {
	return errors.Wrap(w, errOpenDatabaseMsg)
}

func ErrStartHTTPServer(w error) error {
	return errors.Wrap(w, errStartHTTPServerMsg)
}
