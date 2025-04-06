package processes

import (
	"context"
	"io"

	"github.com/pkg/errors"

	"github.com/go-microfrontend/images-s3/internal/storage"
)

type Storage interface {
	GetFile(ctx context.Context, arg storage.GetFileParams) (io.ReadCloser, error)
	PutFile(ctx context.Context, arg storage.PutFileParams) error
}

type Activities struct {
	storage Storage
}

func New(storage Storage) *Activities {
	return &Activities{storage: storage}
}

func (a *Activities) GetImage(ctx context.Context, arg storage.GetFileParams) ([]byte, error) {
	r, err := a.storage.GetFile(ctx, arg)
	if err != nil {
		return nil, errors.Wrap(err, "getting file")
	}
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "reading file")
	}

	return b, nil
}

func (a *Activities) PutImage(ctx context.Context, arg storage.PutFileParams) error {
	return errors.Wrap(a.storage.PutFile(ctx, arg), "putting file")
}
