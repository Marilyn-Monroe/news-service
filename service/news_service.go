package service

import (
	"context"
	"news-service/model"
)

type NewsService interface {
	Update(ctx context.Context, newsModel model.NewsUpdateModel, id int64) model.NewsUpdateModel
	FindAll(ctx context.Context) []model.NewsModel
}
