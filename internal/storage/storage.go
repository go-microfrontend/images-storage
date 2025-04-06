package storage

import (
	"github.com/minio/minio-go/v7"
)

type Storage struct {
	client *minio.Client
}

func New(endpoint string, options *minio.Options) (*Storage, error) {
	client, err := minio.New(endpoint, options)
	if err != nil {
		return nil, err
	}

	return &Storage{
		client: client,
	}, nil
}
