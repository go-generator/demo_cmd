package impl

import (
	. "github.com/common-go/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	. "project_name/model"
	"reflect"
)

type BookCategoryServiceImpl struct {
	database   *mongo.Database
	collection *mongo.Collection
	*DefaultGenericService
	*DefaultSearchService
}

func NewBookCategoryServiceImpl(db *mongo.Database, searchResultBuilder SearchResultBuilder) *BookCategoryServiceImpl {
	var model BookCategory
	modelType := reflect.TypeOf(model)
	collection := "bookCategory"
	mongoService, searchService := NewMongoGenericSearchService(db, modelType, collection, searchResultBuilder, false, "")
	return &BookCategoryServiceImpl{db, db.Collection(collection), mongoService, searchService}
}
