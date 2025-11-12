package repository

import "github.com/luizn/go-url-short/internal/model"

// UrlRepository defines the interface for URL repository operations.
type UrlRepository interface {
	Save(url *model.Url) error
	FindByShortUrl(shortUrl string) (*model.Url, error)
}