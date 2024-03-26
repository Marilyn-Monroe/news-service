package impl

import (
	"context"
	"news-service/entity"
	"news-service/model"
	"news-service/repository"
	"news-service/service"
)

func NewNewsServiceImpl(newsRepository *repository.NewsRepository) service.NewsService {
	return &newsServiceImpl{NewsRepository: *newsRepository}
}

type newsServiceImpl struct {
	repository.NewsRepository
}

func (service *newsServiceImpl) Update(ctx context.Context, newsModel model.NewsUpdateModel, id int64) model.NewsUpdateModel {
	news := entity.News{
		Id:         id,
		Title:      newsModel.Title,
		Content:    newsModel.Content,
		Categories: newsModel.Categories,
	}
	service.NewsRepository.Update(ctx, news)
	return newsModel
}

func (service *newsServiceImpl) FindAll(ctx context.Context) (responses []model.NewsModel) {
	news := service.NewsRepository.FindAll(ctx)
	for _, news := range news {
		responses = append(responses, model.NewsModel{
			Id:         news.Id,
			Title:      news.Title,
			Content:    news.Content,
			Categories: news.Categories,
		})
	}
	if len(news) == 0 {
		return []model.NewsModel{}
	}
	return responses
}
