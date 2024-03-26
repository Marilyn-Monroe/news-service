package impl

import (
	"context"
	"gopkg.in/reform.v1"
	"math"
	"news-service/entity"
	"news-service/exception"
	"news-service/repository"
	"reflect"
)

func NewNewsRepositoryImpl(DB *reform.DB) repository.NewsRepository {
	return &newsRepositoryImpl{DB: DB}
}

type newsRepositoryImpl struct {
	*reform.DB
}

func (repository *newsRepositoryImpl) Update(ctx context.Context, news entity.News) entity.News {
	newsRow, err := repository.DB.WithContext(ctx).SelectOneFrom(entity.NewsTable, "WHERE id = $1", news.Id)
	exception.PanicLogging(err)

	noteExisting := newsRow.(*entity.News)

	if news.Title != "" {
		noteExisting.Title = news.Title
	}

	if news.Content != "" {
		noteExisting.Content = news.Content
	}

	if !reflect.DeepEqual(news.Categories, []int64{math.MinInt64}) {
		noteExisting.Categories = news.Categories

		_, err := repository.DB.WithContext(ctx).DeleteFrom(entity.NewsCategoriesTable, "WHERE newsid = $1", news.Id)
		exception.PanicLogging(err)

		for _, category := range news.Categories {
			categoryModel := &entity.NewsCategories{
				NewsId:     news.Id,
				CategoryId: category,
			}

			err = repository.DB.WithContext(ctx).Save(categoryModel)
			exception.PanicLogging(err)
		}
	}

	err = repository.DB.WithContext(ctx).Save(noteExisting)
	exception.PanicLogging(err)

	return *noteExisting
}

func (repository *newsRepositoryImpl) FindAll(ctx context.Context) []entity.News {
	news, err := repository.DB.WithContext(ctx).SelectAllFrom(entity.NewsTable, "")
	exception.PanicLogging(err)

	res := make([]entity.News, 0, len(news))
	for _, newsStruct := range news {
		note := *newsStruct.(*entity.News)

		categories, err := repository.DB.WithContext(ctx).SelectAllFrom(entity.NewsCategoriesTable, "WHERE NewsId = $1", note.Id)
		exception.PanicLogging(err)

		note.Categories = []int64{}
		for _, category := range categories {
			note.Categories = append(note.Categories, category.(*entity.NewsCategories).CategoryId)
		}

		res = append(res, note)
	}
	return res
}
