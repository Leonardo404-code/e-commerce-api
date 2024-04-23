package google

import (
	"context"
	"mime/multipart"
)

type Google interface {
	Upload(
		ctx context.Context,
		product string,
		file multipart.File,
	) error
}
