package repository

import (
	"context"
	"news-service/entity"
)

type NewsRepository interface {
	Update(ctx context.Context, news entity.News) entity.News
	FindAll(ctx context.Context) []entity.News
}
