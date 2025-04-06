package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/go-microfrontend/images-s3/internal/processes"
	"github.com/go-microfrontend/images-s3/internal/storage"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	client, err := client.Dial(client.Options{HostPort: os.Getenv("TEMPORAL_ADDR")})
	if err != nil {
		slog.Error("failed to connect to temporal", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer client.Close()

	useSSL, err := strconv.ParseBool(os.Getenv("MINIO_USESSL"))
	if err != nil {
		slog.Error("failed to parse MINIO_USESSL", slog.String("error", err.Error()))
		os.Exit(1)
	}

	storage, err := storage.New(os.Getenv("MINIO_ENDPOINT"), &minio.Options{
		Creds: credentials.NewStaticV4(
			os.Getenv("MINIO_ACCESS_KEY"),
			os.Getenv("MINIO_SECRET_KEY"),
			"",
		),
		Secure: useSSL,
	})
	if err != nil {
		slog.Error("failed to connect to minio", slog.String("error", err.Error()))
		os.Exit(1)
	}

	activities := processes.New(storage)

	w := worker.New(client, os.Getenv("TASK_QUEUE"), worker.Options{})
	for _, workflow := range processes.Workflows {
		w.RegisterWorkflow(workflow)
	}
	w.RegisterActivity(activities)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		slog.Error("failed to run worker", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
