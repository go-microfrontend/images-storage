package storage

import (
	"bytes"
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

type GetFileParams struct {
	BucketName string
	ObjectName string
}

func (s *Storage) GetFile(ctx context.Context, arg GetFileParams) (io.ReadCloser, error) {
	obj, err := s.client.GetObject(ctx, arg.BucketName, arg.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "getting object")
	}

	return obj, nil
}

type PutFileParams struct {
	BucketName  string
	ObjectName  string
	Data        []byte
	Size        int64
	ContentType string
}

func (s *Storage) PutFile(ctx context.Context, arg PutFileParams) error {
	err := s.client.MakeBucket(ctx, arg.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		return errors.Wrap(err, "making bucket")
	}

	_, err = s.client.PutObject(
		ctx,
		arg.BucketName,
		arg.ObjectName,
		bytes.NewReader(arg.Data),
		arg.Size,
		minio.PutObjectOptions{ContentType: arg.ContentType},
	)
	if err != nil {
		return errors.Wrap(err, "putting object")
	}

	return nil
}
