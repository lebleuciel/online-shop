package backend

import (
	"github.com/pkg/errors"

	"github.com/lebleuciel/online-shop/internal/backend/server"
)

func NewBackendServer() (*server.Server, error) {
	srv, err := server.NewServer()
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize new backend server")
	}
	return srv, nil
}
