package usercase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/luizn/go-url-short/internal/model"
	"github.com/luizn/go-url-short/internal/repository"
)

type CreateUrlUseCase struct {
	repo repository.UrlRepository
}

func NewCreateUrlUseCase(repo repository.UrlRepository) *CreateUrlUseCase{
	return &CreateUrlUseCase{repo: repo}
}

func (u *CreateUrlUseCase) Execute(ctx context.Context, originalURL string) (*model.Url, error) {
    shortCode := uuid.New().String()[:8] // exemplo simples — pode trocar por algo mais amigável

    url := &model.Url{
        ShortUrl:   shortCode,
        OriginalUrl: originalURL,
        CreatedAt:   time.Now(),
    }

    if err := u.repo.Save(ctx, url); err != nil {
        return nil, err
    }

    return url, nil
}