package processes

import (
	"time"

	"github.com/pkg/errors"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"github.com/go-microfrontend/images-s3/internal/storage"
)

var imageActivityOptions = workflow.ActivityOptions{
	StartToCloseTimeout: time.Minute,
	RetryPolicy: &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    10 * time.Second,
		MaximumAttempts:    5,
	},
}

var Workflows = []any{GetImageWF, PutImageWF}

func GetImageWF(ctx workflow.Context, arg storage.GetFileParams) ([]byte, error) {
	ctx = workflow.WithActivityOptions(ctx, imageActivityOptions)

	var imageBytes []byte
	err := workflow.ExecuteActivity(ctx, "GetImage", arg).Get(ctx, &imageBytes)
	if err != nil {
		return nil, errors.Wrap(err, "executing GetImage activity")
	}

	return imageBytes, nil
}

func PutImageWF(ctx workflow.Context, arg storage.PutFileParams) error {
	ctx = workflow.WithActivityOptions(ctx, imageActivityOptions)

	err := workflow.ExecuteActivity(ctx, "PutImage", arg).Get(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "executing PutImage activity")
	}

	return nil
}
