package config

import (
	"context"
	"project_name/controller"

	"project_name/service/impl"

	"github.com/common-go/mongo"
)

type ApplicationContext struct {
	BookController *controller.BookController

	BookCategoryController *controller.BookCategoryController
}

func NewApplicationContext(mongoConfig mongo.MongoConfig) (*ApplicationContext, error) {
	ctx := context.Background()
	mongoDb, er1 := mongo.SetupMongo(ctx, mongoConfig)
	if er1 != nil {
		return nil, er1
	}

	mongoQueryBuilder := &mongo.DefaultQueryBuilder{}
	mongoSortBuilder := &mongo.DefaultSortBuilder{}
	mongoSearchResultBuilder := &mongo.DefaultSearchResultBuilder{
		Database:     mongoDb,
		QueryBuilder: mongoQueryBuilder,
		SortBuilder:  mongoSortBuilder,
	}
	activityLogService := impl.NewActivityLogServiceImpl(mongoDb)

	bookService := impl.NewBookServiceImpl(mongoDb, mongoSearchResultBuilder)
	bookController := controller.NewBookController(bookService, activityLogService)

	bookCategoryService := impl.NewBookCategoryServiceImpl(mongoDb, mongoSearchResultBuilder)
	bookCategoryController := controller.NewBookCategoryController(bookCategoryService, activityLogService)

	return &ApplicationContext{bookController, bookCategoryController}, nil
}
