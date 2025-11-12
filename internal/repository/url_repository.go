package repository

import (
	"context"

	"github.com/luizn/go-url-short/internal/model"
)

// UrlRepository defines the interface for URL repository operations.
type UrlRepository interface {

	Save(ctx context.Context, url *model.Url) error

	FindByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error)
}