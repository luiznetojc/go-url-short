package repository

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/luizn/go-url-short/internal/model"
)

type UrlRepositoryCassandra struct {
	session *gocql.Session
}

func NewUrlRepositoryCassandra(session *gocql.Session) *UrlRepositoryCassandra {
	return &UrlRepositoryCassandra{session: session}
}

func (r *UrlRepositoryCassandra) Save(ctx context.Context, url *model.Url) error {
	return r.session.Query(`
        INSERT INTO urls (short_code, original_url, created_at)
        VALUES (?, ?, ?)`,
		url.ShortUrl, url.OriginalUrl, url.CreatedAt,
	).WithContext(ctx).Exec()
}

func (r *UrlRepositoryCassandra) FindByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error) {
	query := `SELECT short_url, original_url, created_at FROM urls WHERE short_url = ? LIMIT 1`
	var url model.Url
	if err := r.session.Query(query, shortUrl).WithContext(ctx).Scan(&url.ShortUrl, &url.OriginalUrl, &url.CreatedAt); err != nil {
		return nil, err
	}
	return &url, nil
}
