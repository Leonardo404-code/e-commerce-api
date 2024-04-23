package google

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"strings"
)

func (c *client) Upload(
	ctx context.Context,
	productName string,
	file multipart.File,
) error {
	productNameTrim := strings.Trim(productName, " ")
	productNameReplace := strings.ReplaceAll(productNameTrim, " ", "-")
	productNameToLower := strings.ToLower(productNameReplace)

	writer := c.storage.Bucket(c.bucketName).Object(productNameToLower).NewWriter(ctx)
	defer writer.Close()

	if writer == nil {
		return fmt.Errorf("failed in create a new writer for product %s", productName)
	}

	if _, err := io.Copy(writer, file); err != nil {
		return fmt.Errorf("error in copy file content to writer: %v", err)
	}

	return nil
}
