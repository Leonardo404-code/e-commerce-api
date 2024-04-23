package google

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/Leonargo404-code/e-commerce/internal/env"
	"google.golang.org/api/option"
)

type client struct {
	storage    *storage.Client
	bucketName string
}

func new() (Google, error) {
	credentialsPath := env.GetString(CREDENTIALSPATH)

	storageClient, err := storage.NewClient(
		context.Background(),
		option.WithCredentialsFile(credentialsPath),
	)
	if err != nil {
		return nil, err
	}

	bucketName := env.GetString(BUCKETNAME)

	return &client{
		storage:    storageClient,
		bucketName: bucketName,
	}, nil
}

func Must() Google {
	storage, err := new()
	if err != nil {
		panic(err)
	}

	return storage
}
