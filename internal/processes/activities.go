package processes

import (
	"context"
	"net/url"

	"github.com/pkg/errors"

	"github.com/go-microfrontend/images-storage/internal/storage"
)

type Storage interface {
	GetFile(ctx context.Context, arg storage.GetFileParams) (*url.URL, error)
	PutFile(ctx context.Context, arg storage.PutFileParams) error
}

type Activities struct {
	storage Storage
}

func New(storage Storage) *Activities {
	return &Activities{storage: storage}
}

func (a *Activities) GetImage(ctx context.Context, arg storage.GetFileParams) (string, error) {
	url, err := a.storage.GetFile(ctx, arg)
	if err != nil {
		return "", errors.Wrap(err, "getting file url")
	}

	return url.String(), nil
}

func (a *Activities) PutImage(ctx context.Context, arg storage.PutFileParams) error {
	return errors.Wrap(a.storage.PutFile(ctx, arg), "putting file")
}
